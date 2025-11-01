#!/bin/bash

# Script to run all benchmarks separately, outputting to individual files
# Usage: ./run_benchmarks.sh

set -e

# Create output directory
VERSION=$(git describe --tags --always --dirty)
OUTPUT_DIR="benchmark_results/$VERSION"
mkdir -p "$OUTPUT_DIR"

# Get the directory where the script is located
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

echo "Running benchmarks..."
echo "Results will be saved to: $OUTPUT_DIR/"
echo ""

# List of all benchmark functions
benchmarks=(
    "BenchmarkEncoding_Models"
    "BenchmarkEncoding_TextSizes"
    "BenchmarkEncoding_TextTypes"
    "BenchmarkEncoding_EncodeOrdinary"
    "BenchmarkDecoding"
    "BenchmarkRoundTrip"
    "BenchmarkSpecialTokens"
    "BenchmarkConcurrentEncoding"
    "BenchmarkEncoding_Scaling"
    "BenchmarkEncoding_Allocations"
    "BenchmarkEncoding_File"
    "BenchmarkEncoding_Initialization"
    "BenchmarkEncoding_SpecialTokenConfigs"
    "BenchmarkEncoding_Encodings"
    "BenchmarkEncoding_LargeBatch"
    "BenchmarkConcurrent_EncodeDecode"
    "BenchmarkEncoding_RealisticPatterns"
)

# Function to run a single benchmark
run_benchmark() {
    local benchmark_name=$1
    local output_file="$OUTPUT_DIR/${benchmark_name}.txt"
    local mem_profile="$OUTPUT_DIR/mem.${benchmark_name}.prof"
    local cpu_profile="$OUTPUT_DIR/cpu.${benchmark_name}.prof"
    
    echo "=========================================="
    echo "Running: $benchmark_name"
    echo "Output: $output_file"
    echo "=========================================="
    
    go test -bench="^${benchmark_name}$" -benchmem \
        -memprofile $mem_profile \
        -cpuprofile $cpu_profile \
        -benchtime 2s \
        -count 5 \
        | tee "$output_file"
    
    echo ""
    echo "✓ Completed: $benchmark_name"
    echo ""
}

# Run each benchmark
for benchmark in "${benchmarks[@]}"; do
    run_benchmark "$benchmark"
done

echo "=========================================="
echo "All benchmarks completed!"
echo "Results saved in: $OUTPUT_DIR/"
echo "=========================================="

