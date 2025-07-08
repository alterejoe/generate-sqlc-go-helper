import re

from sqlcgohelper.conversions import DataConversions


class ParseSqlcFile:
    def __init__(self, file):
        self.file = file
        self.parsedqueries = self.parse_out_queries()

    def parse_out_queries(self):
        pattern = re.compile(
            r"--\s*name:\s*([\w]+)\s*:(\w+)\n((?:SELECT|INSERT|UPDATE|DELETE)[^;]*;)",
            re.IGNORECASE,
        )

        data = pattern.findall(self.file)
        sqlc = [ParseSqlcQuery(query) for query in data]
        return sqlc

    def get_raw_params(self, query):
        return re.findall(r"sqlc\.arg\s*\(\s*(\w+)\s*\)", query)

    def get_queries(self):
        return self.parsedqueries


class ParseSqlcQuery:
    def __init__(self, query):
        assert len(query) == 3
        self._name = query[0]
        self._sqlc_return = query[1]
        self._query = query[2]
        self._params = self.get_params()
        self._data = DataConversions(self._name, self._sqlc_return, self._params)

    def get_params(self):
        return re.findall(r"sqlc\.arg\s*\(\s*(\w+)\s*\)", self._query)

    def get_data(self):
        data = {}
        for param in self._params:
            data[param] = self._data
        return self._data
