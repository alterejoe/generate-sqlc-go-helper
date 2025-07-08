import os
import sys
import logging
from sqlcgohelper import parse as Parse
from sqlcgohelper import create as Create
from sqlcgohelper import parse
from sqlcgohelper.parse import ParseSqlcInputFile
from sqlcgohelper.templates import DBQueryFileTemplate, DBQueryParamTemplate


logger = logging.getLogger(__name__)
logger.setLevel(logging.DEBUG)
logger.addHandler(logging.StreamHandler(sys.stdout))


def main(sqlcpath, outputpath):
    files = [
        os.path.join(sqlcpath, f) for f in os.listdir(sqlcpath) if f.endswith(".sql")
    ]
    # parsed_files: []ParseSqlcQuery
    parsed_queries = {}
    for sqlcfile in files:
        if os.path.exists(sqlcfile):
            parsed = ParseSqlcInputFile(open(sqlcfile, "r").read())
            # parsed_queries.extend(parsed.get_queries())
            ## extend parsed_queries
            queries = parsed.get_queries()
            for query in queries:
                parsed_queries[queries[query].get_name()] = queries[query]
        else:
            logger.error(f"File {sqlcfile} does not exist")

    body = ""
    for file in parsed_queries:
        query = parsed_queries[file]
        template = DBQueryParamTemplate(query.get_data())
        rendered = template.render()
        body += rendered

    filetemplate = DBQueryFileTemplate("queries", body)
    file = filetemplate.render()
    print(file)


if __name__ == "__main__":
    args = sys.argv[1:]
    args = ["./sqlc", "./output"]

    sqlcpath = args[0]
    outputpath = args[1]
    main(sqlcpath, outputpath)
