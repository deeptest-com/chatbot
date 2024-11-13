package consts

import "errors"

const (
	ConfigType     = "json"
	CasbinFileName = "rbac_model.conf"
	DirUpload      = "upload"
)

var (
	App = "server"

	ExecDir = ""
	WorkDir = ""

	ConfDir = "config"
)

var (
	DatabaseType = "sqlite"
)

var (
	ErrUserNameOrPassword = errors.New("用户名或密码错误")
	ErrUserNameInvalid    = errors.New("用户名名称已经被使用")
	ErrRoleNameInvalid    = errors.New("角色名称已经被使用")

	ErrParamValidate      = errors.New("参数验证失败")
	ErrPaginateParam      = errors.New("分页查询参数缺失")
	ErrUnSupportFramework = errors.New("不支持的框架")
)

var (
	Instructions = `
[
    {"name": ”greetings", "steps": []},
    {"name": ”confirm", "steps": []},
    
    {"name": "create_part", "steps": ["init", "input_part_no", "input_part_name", "show_part_form"]},
    {"name": "attach_material", "steps": ["init", "input_materials", "fill_material_form"]},
    {"name": "attach_geometry", "steps": ["init", "input_geometry", "input_geometry_version", "input_design", "input_design_version", "input_drawing", "input_drawing_version"]},
    {"name": "create_structure", "steps": ["init", "fill_st_form"]},
    {"name": "assign_project", "steps": ["init", "input_project"]},
    {"name": "data_check", "steps": ["init"]},
    {"name": "freeze_structure", "steps": ["init", "freeze_confirm", "input_design", "input_drawing", "input_geometry_version"]},
    {"name": "submit_structure", "steps": ["init", "submit_st_confirm", "fix_st"]},
    {"name": "track_st", "steps": ["init"]
]
`
)
