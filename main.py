import os
import sys
import logging
from sqlcgohelper import parse as Parse
from sqlcgohelper import create as Create


logger = logging.getLogger(__name__)
logger.setLevel(logging.DEBUG)
logger.addHandler(logging.StreamHandler(sys.stdout))


def main(sqlcpath, outputpath):
    files = [
        os.path.join(sqlcpath, f) for f in os.listdir(sqlcpath) if f.endswith(".sql")
    ]
    sqlc = {}
    for sqlcfile in files:
        if os.path.exists(sqlcfile):
            file = Parse.parse_sqlc_file(open(sqlcfile, "r").read())
            assert file is not None
            sqlc.update(file)  ## havent seen this update function before but it works

            # logger.info(f"File {sqlcfile} loaded")
        else:
            logger.error(f"File {sqlcfile} does not exist")

    rendered_structs = []
    for queryname in sqlc:
        rendered_struct = Create.create_param_struct(queryname, sqlc[queryname])
        rendered_structs.append(rendered_struct)

    rendered_file = Create.create_query_file("queries", rendered_structs)
    # print(rendered_file)

    # Parse.parse_sqla()


# type GetLastUpdated struct {
# }
#
# func (glu GetLastUpdated) Query(query *db.Queries, r context.Context) (any, error) {
# 	// this will error to a default pgtype.Timestamptz which is valid and very stale
# 	lastupdated, _ := query.GetLastUpdated(r)
#
# 	return lastupdated, nil
# }
# type BatchPostAccountData struct {
# 	Params *db.BatchPostAccountDataParams
# }
#
# func (bpd BatchPostAccountData) Query(query *db.Queries, r context.Context) (any, error) {
# 	if bpd.Params == nil {
# 		return nil, nil
# 	}
# 	err := query.BatchPostAccountData(r, *bpd.Params)
#
# 	return nil, err
# }

if __name__ == "__main__":
    args = sys.argv[1:]
    args = ["./sqlc", "./output"]
    # if len(args) == 0:
    #     args = ["./sqlc"]

    sqlcpath = args[0]
    outputpath = args[1]
    main(sqlcpath, outputpath)
