"""Cost estimate for a session, in USD, from its usage totals."""

RATE_INPUT_PER_1K = 0.003
RATE_OUTPUT_PER_1K = 0.015
CACHE_READ_DISCOUNT = 0.1


def session_cost(summary):
    input_tokens = summary["input_tokens"]
    output_tokens = summary["output_tokens"]
    return (input_tokens / 1000) * RATE_INPUT_PER_1K + (output_tokens / 1000) * RATE_OUTPUT_PER_1K
