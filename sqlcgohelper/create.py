from jinja2 import Environment, FileSystemLoader

env = Environment(loader=FileSystemLoader("./sqlcgohelper/jinja2"))


def get_capital_letters(input):
    chars = "".join([char for char in input if char.isupper()])
    return chars


special = ["id"]


def sqlc_column_name(column):
    underscore_split = column.split("_")
    capitalized = []
    for i, word in enumerate(underscore_split):
        if i == 0:
            capitalized.append(word)
        elif word in special:
            capitalized.append(word.upper())
        else:
            capitalized.append(word.capitalize())
    word = "".join(capitalized)
    return word


def create_param_struct(sqlcname, params):
    """
    Params:
        - name
        - sqlc_return
        - sqlc_params
        - query
    """
    print("Name: ", sqlcname)
    print("\tSQLC params: ", params["sqlc_params"])

    short_name = get_capital_letters(sqlcname).lower()
    lower_name = sqlcname.lower()
    template = env.get_template("params.jinja")
    return template.render(
        sqlcname=sqlcname,
        short_sqlcname=short_name,
        lower_sqlcname=lower_name,
    )


def create_query_file(package_name, renderedqueries=[]):
    query_text = "\n".join([query for query in renderedqueries])
    template = env.get_template("query-file.jinja")
    return template.render(package_name=package_name, body=query_text)
