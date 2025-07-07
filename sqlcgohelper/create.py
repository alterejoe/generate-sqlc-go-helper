from jinja2 import Environment, FileSystemLoader

env = Environment(loader=FileSystemLoader("./sqlcgohelper/jinja2"))


def get_capital_letters(input):
    chars = "".join([char for char in input if char.isupper()])
    return chars


def create_param_struct(sqlcname, params):
    short_name = get_capital_letters(sqlcname).lower()
    lower_name = sqlcname.lower()
    template = env.get_template("params.jinja")
    return template.render(
        sqlcname=sqlcname,
        short_sqlcname=short_name,
        lower_sqlcname=lower_name,
    )
