from jinja2 import Environment, FileSystemLoader, Template

env = Environment(loader=FileSystemLoader("./sqlcgohelper/jinja2"))


def get_capital_letters(input):
    chars = "".join([char for char in input if char.isupper()])
    return chars


def clean_params(sqlcparams, sqlcname):
    if len(sqlcparams) > 1:
        return sqlcname + "Params"
    elif len(sqlcparams) == 1:
        return sqlcparams[0]
    else:
        return ""


returns = ["one", "many"]


values = {
    "sqlc_function": None,  ## sqlc name
    "struct_params": None,  ## None | Params (*db.{{ sqlc_function }}Params  | *{{ one_param_name }}
    "sqlc_abbv": None,
    "if_params": None,
    "query_return": None,
    "query_params": None,
    "return_value": None,
}


def create_param_struct(sqlcname, params):
    queryptr_template = env.get_template("query-pointer.jinja")
    ifparams_template = env.get_template("if-params.jinja")


def create_param_struct_old(sqlcname, params):
    """
    Params:
    { name sqlc_return sqlc_params params query }
    """
    print("Name: ", sqlcname)
    print("\tSQLC params: ", params["sqlc_params"])

    short_name = get_capital_letters(sqlcname).lower()
    lower_name = sqlcname.lower()
    lower_err = ""
    if params["sqlc_return"] in returns:
        lower_err = ", err"

    template = env.get_template("params.jinja")
    cleanparam = clean_params(params["sqlc_params"], sqlcname)
    commaparam = "r"

    if "Params" in cleanparam:
        commaparam = f"r, *{short_name}.Params"
    elif len(cleanparam) > 0:
        commaparam = f"r, *{short_name}.{cleanparam}"

    structparam = ""
    if cleanparam != "":
        if "Params" in cleanparam:
            structparam = f"Params *db.{cleanparam}"
        else:
            structparam = f"{cleanparam} SOMETYPE"
    return template.render(
        sqlcname=sqlcname,
        short_sqlcname=short_name,
        lower_sqlcname=lower_name,
        struct_sqlcparams=structparam,
        comma_sqlcparams=commaparam,
    )


def create_query_file(package_name, renderedqueries=[]):
    query_text = "\n".join([query for query in renderedqueries])
    template = env.get_template("query-file.jinja")
    return template.render(package_name=package_name, body=query_text)
