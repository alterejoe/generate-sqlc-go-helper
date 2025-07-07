import os
import sys
import logging
from sqlcgohelper import parse as Parse
from sqlcgohelper import create as Create


logger = logging.getLogger(__name__)
logger.setLevel(logging.DEBUG)
logger.addHandler(logging.StreamHandler(sys.stdout))


def main(path):
    files = [os.path.join(path, f) for f in os.listdir(path) if f.endswith(".sql")]
    sqlc = {}
    for sqlcfile in files:
        if os.path.exists(sqlcfile):
            file = Parse.parse_sqlc_file(open(sqlcfile, "r").read())
            assert file is not None
            sqlc.update(file)  ## havent seen this update function before but it works

            logger.info(f"File {sqlcfile} loaded")
        else:
            logger.error(f"File {sqlcfile} does not exist")

    for queryname in sqlc:
        print(queryname)
        renderedparamstruct = Create.create_param_struct(
            queryname, sqlc[queryname]["params"]
        )
        print(renderedparamstruct)

    Parse.parse_sqla()


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
    if len(args) == 0:
        args = ["./sqlc"]
    path = args[0]
    main(path)
