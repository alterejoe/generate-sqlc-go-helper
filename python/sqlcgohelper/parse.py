import logging
import sys
from abc import ABC, abstractmethod
import os
import re

from sqlcgohelper.data import Data

logger = logging.getLogger(__name__)
logger.setLevel(logging.DEBUG)
logger.addHandler(logging.StreamHandler(sys.stdout))


class Sqlc:
    def __init__(self):
        self._header = None  ## -- name: Title :operation
        self._query = None  ## full sql query


class SqlcImporter:
    def __init__(self):
        self._rawfiledata = {}
        self._functiondata = {}  ## using the function name as the key

    def add_function_data(self, name, data):
        currentdata = self._functiondata.get(name)
        if currentdata:
            for key in data:
                if key in currentdata:
                    previous = currentdata[key]
                    current = data[key]
                    logger.warning(f"Duplicate key {key} in function {name}")
                    logger.warning(f"{previous} => {current}")
                else:
                    currentdata[key] = data[key]

    def get_function(self, name):
        return self._functiondata.get(name)

    def export(self):
        return self._functiondata


class SqlcDirImporter(SqlcImporter):
    def __init__(self, dir, pattern=""):
        super().__init__()

        assert os.path.exists(dir)
        self._dir = dir
        self._file_paths = [f for f in os.listdir(dir) if f.endswith(pattern)]

    def _import(self):
        for path in self._file_paths:
            assert os.path.exists(path)
            with open(path, "r") as f:
                self._rawfiledata[path] = f.read()


class SqlcFileImporter(SqlcImporter):
    def __init__(self, filepath, pattern=""):
        assert os.path.exists(filepath)
        assert filepath.endswith(pattern)
        self._filepath = filepath

    def _import(self):
        assert os.path.exists(self._filepath)
        self._rawfiledata = {}
        with open(self._filepath, "r") as f:
            self._rawfiledata[self._filepath] = f.read()
        return self._rawfiledata


class SqlcInput(SqlcDirImporter):
    def __init__(self, dir):
        super().__init__(dir, ".sql")


class SqlcQueryOutput(SqlcDirImporter):
    def __init__(self, dir):
        super().__init__(dir, ".sql.go")


class SqlcModelOutput(SqlcFileImporter):
    def __init__(self, filepath):
        super().__init__(filepath)
