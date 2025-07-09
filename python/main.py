import os
import sys
import logging
from sqlcgohelper import parse as Parse
from sqlcgohelper import create as Create
from sqlcgohelper import parse
from sqlcgohelper.parse import SqlcInput, SqlcModelOutput, SqlcQueryOutput
from sqlcgohelper.templates import DBQueryFileTemplate, DBQueryParamTemplate


logger = logging.getLogger(__name__)
logger.setLevel(logging.DEBUG)
logger.addHandler(logging.StreamHandler(sys.stdout))


paths = {
    "sqlc_input": "./sqlc/",
    "sqlc_output_models": "/home/altjoe/Documents/projects/budget/web-budget/web/db/models.go",
    "sqlc_output_queries": "/home/altjoe/Documents/projects/budget/web-budget/web/db/",
}


def main():

    output_queries = SqlcInput(paths["sqlc_output_queries"])
    output_models = SqlcModelOutput(paths["sqlc_output_models"])
    input_queries = SqlcQueryOutput(paths["sqlc_input"])


# def old_main(sqlcpath, outputpath):
#     files = [
#         os.path.join(sqlcpath, f) for f in os.listdir(sqlcpath) if f.endswith(".sql")
#     ]
#     # parsed_files: []ParseSqlcQuery
#     parsed_queries = {}
#     for sqlcfile in files:
#         if os.path.exists(sqlcfile):
#             parsed = ParseSqlcInputFile(open(sqlcfile, "r").read())
#             # parsed_queries.extend(parsed.get_queries())
#             ## extend parsed_queries
#             queries = parsed.get_queries()
#             for query in queries:
#                 parsed_queries[queries[query].get_name()] = queries[query]
#         else:
#             logger.error(f"File {sqlcfile} does not exist")
#
#     body = ""
#     for file in parsed_queries:
#         query = parsed_queries[file]
#         template = DBQueryParamTemplate(query.get_data())
#         rendered = template.render()
#         body += rendered
#
#     filetemplate = DBQueryFileTemplate("queries", body)
#     file = filetemplate.render()
#     print(file)
#

if __name__ == "__main__":
    args = sys.argv[1:]
    args = ["./sqlc", "./output"]

    sqlcpath = args[0]
    outputpath = args[1]
    main()
