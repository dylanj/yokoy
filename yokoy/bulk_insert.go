package yokoy

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/aarondl/null/v9"
)

const bulkInsertBatchSize = 1000

// BulkInsertOptions provides options for bulk insert operations
type BulkInsertOptions struct {
	// OnConflict specifies the conflict resolution strategy
	// If empty, no conflict handling is used (will fail on duplicates)
	// Use "DO NOTHING" to ignore conflicts
	// Use "DO UPDATE SET ..." for update on conflict
	OnConflict string
	// ConflictTarget specifies the conflict target columns (for ON CONFLICT)
	// If empty and OnConflict is set, uses primary key constraint
	ConflictTarget string
}

// BulkInsert performs a bulk insert of records into a PostgreSQL table.
// tableName: name of the table
// columns: column names in order
// values: slice of value slices, where each inner slice contains values for one row
// options: optional bulk insert options for conflict handling
func BulkInsert(ctx context.Context, db *sql.DB, tableName string, columns []string, values [][]interface{}, options ...BulkInsertOptions) error {
	if len(values) == 0 {
		return nil
	}

	// Validate inputs
	if len(columns) == 0 {
		return fmt.Errorf("bulk insert: no columns specified for table %s", tableName)
	}

	// Check that all rows have the same number of values as columns
	for i, row := range values {
		if len(row) != len(columns) {
			return fmt.Errorf("bulk insert: row %d has %d values but expected %d columns for table %s", i, len(row), len(columns), tableName)
		}
	}

	var opts BulkInsertOptions
	if len(options) > 0 {
		opts = options[0]
	}

	// Process in batches to avoid memory issues and parameter limits
	totalBatches := (len(values) + bulkInsertBatchSize - 1) / bulkInsertBatchSize
	for i := 0; i < len(values); i += bulkInsertBatchSize {
		end := i + bulkInsertBatchSize
		if end > len(values) {
			end = len(values)
		}

		batch := values[i:end]
		batchNum := (i / bulkInsertBatchSize) + 1
		if err := bulkInsertBatch(ctx, db, tableName, columns, batch, opts, batchNum, totalBatches); err != nil {
			return fmt.Errorf("bulk insert batch %d/%d failed at offset %d for table %s: %w", batchNum, totalBatches, i, tableName, err)
		}
	}

	return nil
}

func bulkInsertBatch(ctx context.Context, db *sql.DB, tableName string, columns []string, values [][]interface{}, opts BulkInsertOptions, batchNum, totalBatches int) error {
	if len(values) == 0 {
		return nil
	}

	// Build the query
	quotedColumns := make([]string, len(columns))
	for i, col := range columns {
		quotedColumns[i] = fmt.Sprintf(`"%s"`, col)
	}

	// Build VALUES clause with placeholders
	placeholders := make([]string, len(values))
	args := make([]interface{}, 0, len(values)*len(columns))
	argIndex := 1

	for i, row := range values {
		rowPlaceholders := make([]string, len(columns))
		for j := range columns {
			rowPlaceholders[j] = fmt.Sprintf("$%d", argIndex)
			args = append(args, row[j])
			argIndex++
		}
		placeholders[i] = fmt.Sprintf("(%s)", strings.Join(rowPlaceholders, ", "))
	}

	// Build the base INSERT query
	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES %s",
		tableName,
		strings.Join(quotedColumns, ", "),
		strings.Join(placeholders, ", "),
	)

	// Add ON CONFLICT clause if specified
	if opts.OnConflict != "" {
		if opts.ConflictTarget != "" {
			query += fmt.Sprintf(" ON CONFLICT (%s) %s", opts.ConflictTarget, opts.OnConflict)
		} else {
			// Default to primary key constraint
			query += fmt.Sprintf(" ON CONFLICT %s", opts.OnConflict)
		}
	}

	_, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("execution failed for batch %d/%d: %w", batchNum, totalBatches, err)
	}

	return nil
}

// NullStringValue converts a null.String to interface{} for bulk insert
func NullStringValue(ns null.String) interface{} {
	if ns.Valid {
		return ns.String
	}
	return nil
}

// NullBoolValue converts a null.Bool to interface{} for bulk insert
func NullBoolValue(nb null.Bool) interface{} {
	if nb.Valid {
		return nb.Bool
	}
	return nil
}

// NullIntValue converts a null.Int to interface{} for bulk insert
func NullIntValue(ni null.Int) interface{} {
	if ni.Valid {
		return ni.Int
	}
	return nil
}

// NullTimeValue converts a null.Time to interface{} for bulk insert
func NullTimeValue(nt null.Time) interface{} {
	if nt.Valid {
		return nt.Time
	}
	return nil
}

