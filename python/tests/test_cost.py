import os
import sys

sys.path.insert(0, os.path.join(os.path.dirname(__file__), ".."))

from cost import session_cost
from sessions import summarize


def test_cached_reads_bill_at_the_discounted_rate():
    # Expected value follows README.md's pricing table.
    summary = summarize("sess-cache-heavy")
    assert round(session_cost(summary), 4) == 0.0109
