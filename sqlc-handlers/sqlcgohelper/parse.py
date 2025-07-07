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
        print(len(query))
        querydata = {
            "name": query[0],
            "params": query[1],
            "query": query[2],
        }
        queries[querydata["name"]] = querydata

    return queries


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
