#!/bin/bash

# Script to compare benchmark results between baseline and patch directories using benchstat
# Usage: ./compare_benchmarks.sh [baseline_dir] [patch_dir]

set -e

# Check if benchstat is available
if ! command -v benchstat &> /dev/null; then
    echo "Error: benchstat not found. Install it with:"
    echo "  go install golang.org/x/perf/cmd/benchstat@latest"
    exit 1
fi

BASELINE_DIR="${1:-benchmark_results/baseline}"
PATCH_DIR="${2:-benchmark_results/patch}"

if [ ! -d "$BASELINE_DIR" ]; then
    echo "Error: Baseline directory not found: $BASELINE_DIR"
    echo "Usage: $0 [baseline_dir] [patch_dir]"
    exit 1
fi

if [ ! -d "$PATCH_DIR" ]; then
    echo "Error: Patch directory not found: $PATCH_DIR"
    echo "Usage: $0 [baseline_dir] [patch_dir]"
    exit 1
fi

echo "Comparing benchmarks with benchstat:"
echo "  Baseline: $BASELINE_DIR"
echo "  Patch:    $PATCH_DIR"
echo ""

# Find all benchmark files in baseline
for baseline_file in "$BASELINE_DIR"/*.txt; do
    if [ ! -f "$baseline_file" ]; then
        continue
    fi
    
    filename=$(basename "$baseline_file")
    patch_file="$PATCH_DIR/$filename"
    
    if [ ! -f "$patch_file" ]; then
        echo "⚠ Missing in patch: $filename"
        continue
    fi
    
    # Skip if files are empty
    if [ ! -s "$baseline_file" ] || [ ! -s "$patch_file" ]; then
        echo "⚠ Skipping $filename (one or both files are empty)"
        continue
    fi
    
    echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
    echo "📊 $filename"
    echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
    benchstat "$baseline_file" "$patch_file" 2>&1 || {
        exit_code=$?
        if [ $exit_code -ne 0 ]; then
            echo "⚠ benchstat exited with code $exit_code for $filename"
        fi
    }
    echo ""
done

echo "✓ Comparison complete"

