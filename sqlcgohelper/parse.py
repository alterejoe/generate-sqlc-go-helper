import os
import re

from sqlcgohelper.conversions import Data


class ParseSqlcInputFile:
    def __init__(self, file):
        self.file = file
        self.parsedqueries = self.parse_out_queries()

    def parse_out_queries(self):
        pattern = re.compile(
            r"--\s*name:\s*([\w]+)\s*:(\w+)\n((?:SELECT|INSERT|UPDATE|DELETE)[^;]*;)",
            re.IGNORECASE,
        )

        data = pattern.findall(self.file)
        # sqlc = [ParseSqlcInputQuery(query) for query in data]
        sqlc = {}
        for query in data:
            parsed = ParseSqlcInputQuery(query)
            sqlc[parsed.get_name()] = parsed

        return sqlc

    def get_raw_params(self, query):
        return re.findall(r"sqlc\.arg\s*\(\s*(\w+)\s*\)", query)

    def get_queries(self):
        return self.parsedqueries


class ParseSqlcOutputQueries:
    def __init__(self, path):
        if path.endswith("/"):
            raise ValueError
        self.path = path

    def import_queries(self):
        files = [f for f in os.listdir(self.path) if f.endswith(".sql.go")]
        for file in files:
            data = ParseSqlcOutputModels(os.path.join(self.path, file)).import_models()
            pass


class ParseSqlcOutputModels:
    def __init__(self, path):
        if path.endswith("models.go"):
            raise ValueError

        self.path = path

    def import_models(self):
        pass


class ParseSqlcInputQuery:
    def __init__(self, query):  ## coupling here
        assert len(query) == 3
        self._name = query[0]
        self._sqlc_return = query[1]
        self._query = query[2]
        self._params = self.get_params()
        self._data = Data(self._name, self._sqlc_return, self._params)

    def get_params(self):
        return re.findall(r"sqlc\.arg\s*\(\s*(\w+)\s*\)", self._query)

    def get_data(self):
        data = {}
        for param in self._params:
            data[param] = self._data
        return self._data

    def get_name(self):
        return self._name
