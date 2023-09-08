import sys
from argparse import Namespace

import requests

from determined import cli
from determined.common.declarative_argparse import Arg, ArgsDescription, Cmd


# Print the body of a response in chunks so we don't have to buffer the whole thing.
def print_response(r: requests.Response) -> None:
    for chunk in r.iter_content(chunk_size=4096):
        sys.stdout.buffer.write(chunk)


def raw(args: Namespace) -> None:
    sess = cli.setup_session(args)
    params = {"timestamp_after": args.timestamp_after, "timestamp_before": args.timestamp_before}
    path = "api/v1/resources/allocation/raw" if args.json else "resources/allocation/raw"
    print_response(sess.get(path, params=params))


def aggregated(args: Namespace) -> None:
    sess = cli.setup_session(args)
    params = {
        "start_date": args.start_date,
        "end_date": args.end_date,
        "period": "RESOURCE_ALLOCATION_AGGREGATION_PERIOD_MONTHLY"
        if args.monthly
        else "RESOURCE_ALLOCATION_AGGREGATION_PERIOD_DAILY",
    }
    path = (
        "api/v1/resources/allocation/aggregated" if args.json else "resources/allocation/aggregated"
    )
    print_response(sess.get(path, params=params))


args_description: ArgsDescription = [
    Cmd(
        "res|ources",
        None,
        "query historical resource allocation",
        [
            Cmd(
                "raw",
                raw,
                "get raw allocation information",
                [
                    Arg("timestamp_after"),
                    Arg("timestamp_before"),
                    Arg("--json", action="store_true", help="output JSON rather than CSV"),
                ],
            ),
            Cmd(
                "agg|regated",
                aggregated,
                "get aggregated allocation information",
                [
                    Arg("start_date", help="first date to include"),
                    Arg("end_date", help="last date to include"),
                    Arg("--json", action="store_true", help="output JSON rather than CSV"),
                    Arg(
                        "--monthly",
                        action="store_true",
                        help="aggregate by month rather than by day",
                    ),
                ],
            ),
        ],
    )
]
