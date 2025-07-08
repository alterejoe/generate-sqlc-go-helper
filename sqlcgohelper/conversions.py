class DataConversions:
    def __init__(self, functionname, sqlc_return, sqlc_params):
        """
        functionname: str
        sqlc_return: str "one, many, exec"
        sqlc_params: list[str] - parameters defined within query `sqlc.arg(param)`
        """
        self._sqlc_return = sqlc_return
        self._sqlc_params = sqlc_params
        self._sqlc_capital_params = [
            self.capitalize_sqlc_column(c) for c in self._sqlc_params
        ]

        self.name = functionname
        self.structparams = self.get_structparams()
        self.lowername = self.get_lowername()
        self.lowererror = self.get_lowererror()
        self.abbv = self.get_abbv()
        self.queryreturn = self.get_queryreturn()
        self.queryparams = self.get_queryparams()
        self.returnvalue = self.get_returnvalue()

    def get_abbv(self):
        input = self.name
        chars = "".join([char for char in input if char.isupper()])
        return chars.lower()

    def get_lowername(self):
        return self.name.lower()

    def get_lowererror(self):
        if self._sqlc_return in ["one", "many"]:
            return ", err"
        else:
            return ", nil"

    def get_structparams(self):
        if len(self._sqlc_params) > 1:
            return f"Params *db.{self.name}Params"
        elif len(self._sqlc_params) == 1:
            return f"{self._sqlc_capital_params[0]} SOMETYPE"
        return ""

    def get_queryparams(self):
        if len(self._sqlc_params) > 1:
            return f"r, *db.{self.name}Params"
        elif len(self._sqlc_params) == 1:

    def get_queryreturn(self):
        return f"{self.lowername}{self.lowererror}"

    def get_returnvalue(self):
        return f"{self.lowername}{self.lowererror}"

    def capitalize_sqlc_column(self, column, solo=False):
        special = ["id"]
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
