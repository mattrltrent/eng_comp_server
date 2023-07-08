from typing import List
import requests
from urllib import parse



class BannerClient(requests.Session):
    def __init__(self):
        super().__init__()
        self.banner_base = "https://banner.uvic.ca/StudentRegistrationSsb/ssb/"
        self.banner_set_term = parse.urljoin(
            self.banner_base, "term/search?mode=search"
        )
        self.banner_get_term = parse.urljoin(
            self.banner_base, "classSearch/getTerms?offset=1&max=1"
        )
        self.banner_search_result = parse.urljoin(
            self.banner_base, "searchResults/searchResults"
        )
        self.term = "202305"  # NOTE: set by self.get_latest_term

        self._MAX_SIZE = 500  # NOTE: max from banner is 500


    def get_data(self, offset: int, term: str = "202305") -> list | None:
        self.term = self.term or term
        q = {
            "txt_term": self.term,
            "pageOffset": offset * self._MAX_SIZE,
            "pageMaxSize": self._MAX_SIZE,
        }
        url = f"{self.banner_search_result}?{parse.urlencode(q)}"
        print(url)
        resp = self.get(url)
        assert resp.status_code == 200

        payload = resp.json()

        data = payload["data"]
        if len(data) == 0:
            print("no data found")
            return None

        fetched = offset * self._MAX_SIZE + len(data)
        print(f"fetched {fetched}/{payload['sectionsFetchedCount']}")
        return data