from sqlcgohelper._models import Base
import re
import logging
import sys
import json


logger = logging.getLogger(__name__)
logger.setLevel(logging.DEBUG)
logger.addHandler(logging.StreamHandler(sys.stdout))


def parse_sqlc_file(file):
    # pattern -
    pattern = re.compile(
        r"--\s*name:\s*([\w]+)\s*:(\w+)\n((?:SELECT|INSERT|UPDATE|DELETE)[^;]*;)",
        re.IGNORECASE,
    )

    data = pattern.findall(file)
    queries = {}
    for query in data:
        querydata = {
            "name": query[0],
            "sqlc_return": query[1],
            "sqlc_params": parse_sqlc_query(query[2]),
            "query": query[2],
        }
        queries[querydata["name"]] = querydata

    return queries


special = ["id"]


def sqlc_column_name(column, solo=False):
    underscore_split = column.split("_")
    capitalized = []

    for i, word in enumerate(underscore_split):
        if i == 0 and solo:
            capitalized.append(word)
        elif word in special:
            capitalized.append(word.upper())
        else:
            capitalized.append(word.capitalize())
    word = "".join(capitalized)
    return word


def parse_sqlc_query(query):
    params = []
    ## pattern "sqlc.args(paramkey)
    query = query.replace("sqlc.arg (", "sqlc.arg(")
    matches = re.findall(r"sqlc.arg\((\w+)\)", query)
    if len(matches) > 1:
        for match in matches:
            params.append(sqlc_column_name(match))
    elif len(matches) == 1:
        params.append(sqlc_column_name(matches[0], solo=True))

    return params


## this is to eventually build sqlc queries from tables I've created
def parse_sqla():
    for table in Base.metadata.tables:
        print(table)


"""
What i need to do from here:

take sqlc name and create a struct

create a query function for that struct that has necissary format

check to see if it has parameters

if it does add Params to struct and find correct generated struct to fill

i think if it is multiple params it gets <sqlcname>Params 

if it is a single param it gets the name of the param


"""
