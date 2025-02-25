import time
from typing import Optional

import requests

from determined.common import api
from determined.common.api import certs

from .errors import MasterTimeoutExpired

DEFAULT_TIMEOUT = 100


def _make_master_url(master_host: str, master_port: int, suffix: str = "") -> str:
    return "http://{}:{}/{}".format(master_host, master_port, suffix)


def wait_for_master(
    master_host: str,
    master_port: int = 8080,
    timeout: int = DEFAULT_TIMEOUT,
    cert: Optional[certs.Cert] = None,
) -> None:
    master_url = _make_master_url(master_host, master_port)

    return wait_for_master_url(master_url, timeout, cert)


def wait_for_master_url(
    master_url: str,
    timeout: int = DEFAULT_TIMEOUT,
    cert: Optional[certs.Cert] = None,
) -> None:
    POLL_INTERVAL = 2
    polling = False
    start_time = time.time()

    try:
        while time.time() - start_time < timeout:
            try:
                r = api.get(master_url, "info", authenticated=False, cert=cert)
                if r.status_code == requests.codes.ok:
                    return
            except api.errors.MasterNotFoundException:
                pass
            if not polling:
                polling = True
                print("Waiting for master instance to be available...", end="", flush=True)
            time.sleep(POLL_INTERVAL)
            print(".", end="", flush=True)

        raise MasterTimeoutExpired
    finally:
        if polling:
            print()


def wait_for_genai_url(
    master_url: str,
    timeout: int = DEFAULT_TIMEOUT,
    cert: Optional[certs.Cert] = None,
) -> None:
    POLL_INTERVAL = 2
    polling = False
    start_time = time.time()
    GENAI_PREFIX = "/genai"
    check_path = GENAI_PREFIX + "/api/v1/workspaces"

    try:
        while time.time() - start_time < timeout:
            try:
                auth = api.Authentication(master_address=master_url, cert=cert)
                r = api.get(master_url, check_path, authenticated=True, cert=cert, auth=auth)
                print(r)
                print(r.status_code)
                if r.status_code == requests.codes.ok:
                    _ = r.json()
                    return
            except (api.errors.MasterNotFoundException, api.errors.APIException):
                pass
            if not polling:
                polling = True
                print("Waiting for GenAI instance to be available...", end="", flush=True)
            time.sleep(POLL_INTERVAL)
            print(".", end="", flush=True)
        raise MasterTimeoutExpired
    finally:
        if polling:
            print()
