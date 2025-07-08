from jinja2 import Template
from abc import ABC, abstractmethod
from sqlcgohelper.conversions import DataConversions


class Jinja(ABC):
    def __init__(self):
        self.template = self.get_template()

    def render(self):
        assert self.template is not None
        assert isinstance(self.template, str)
        template = Template(self.template)
        return template.render(**self.get_params())

    @abstractmethod
    def get_template(self) -> str:
        pass

    @abstractmethod
    def get_params(self) -> dict:
        pass


class DBQueryParamTemplate(Jinja):
    def __init__(self, data: DataConversions):
        super().__init__()
        self.data = data

    def get_template(self):
        return """type {{ name }} struct {
    {{ structparams | default("")  }}
}

func ({{ abbv }} {{ name }}) Query(query *db.Queries, r context.Context) (any, error) {
        {{ ifparams }}
        {{ queryreturn  }} := query.{{ name }}({{ queryparams }})
        return {{ returnvalue }}
}
        """

    def get_params(self):
        return {
            "name": self.data.name,
            "structparams": self.data.structparams,
            "abbv": self.data.abbv,
            "lowername": self.data.lowername,
            "lowererror": self.data.lowererror,
            "queryreturn": self.data.queryreturn,
            "queryparams": self.data.queryparams,
            # "ifparams": DBQueryIfParamTemplate(self.data).render(),
            "ifparams": (
                ""
                if len(self.data._sqlc_params) == 0
                else DBQueryIfParamTemplate(self.data.abbv).render()
            ),
            "returnvalue": self.data.returnvalue,
        }


class DBQueryIfParamTemplate(Jinja):
    def __init__(self, abbv):
        super().__init__()
        self.abbv = abbv

    def get_template(self):
        return """if {{ abbv }}.Params == nil {
    return nil, errors.New("params is nil")
}
        """

    def get_params(self):
        return {
            "abbv": self.abbv,
        }


class DBQueryFileTemplate(Jinja):
    def __init__(self, package_name, body):
        super().__init__()
        self.package_name = package_name
        self.body = body

    def get_template(self):
        return """package {{ package_name }}
import (
	"context"
	"github.com/alterejoe/budget/db"
    "errors"
)

{{ body }}
        """

    def get_params(self):
        return {
            "package_name": self.package_name,
            "body": self.body,
        }
