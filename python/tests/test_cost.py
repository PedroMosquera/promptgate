import os
import sys

sys.path.insert(0, os.path.join(os.path.dirname(__file__), ".."))

from cost import session_cost
from sessions import summarize


def test_cached_reads_bill_at_the_discounted_rate():
    summary = summarize("sess-cache-heavy")
    # 90 fresh input tokens at full rate, 8310 cached at the 10% rate,
    # 540 output tokens at the output rate. See README.md's pricing table.
    assert round(session_cost(summary), 4) == 0.0109
