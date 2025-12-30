#!/bin/bash
# IDRM 项目初始化脚本
# 用法: ./scripts/init-project.sh [options]
#
# 选项:
#   -n, --name      项目名称 (默认: my-project)
#   -m, --module    Go 模块路径 (默认: github.com/example/<project_name>)
#   --with-rpc      包含 RPC 服务
#   --with-job      包含 Job 服务
#   --with-consumer 包含 Consumer 服务
#   -a, --all       包含所有可选服务
#   -i, --interactive 交互式模式
#   -h, --help      显示帮助

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 默认值
PROJECT_NAME="my-project"
MODULE_PATH=""
WITH_RPC=false
WITH_JOB=false
WITH_CONSUMER=false
INTERACTIVE=false

# 显示帮助
show_help() {
    echo "IDRM 项目初始化脚本"
    echo ""
    echo "用法: ./scripts/init-project.sh [options]"
    echo ""
    echo "选项:"
    echo "  -n, --name      项目名称 (默认: my-project)"
    echo "  -m, --module    Go 模块路径 (默认: github.com/example/<project_name>)"
    echo "  --with-rpc      包含 RPC 服务"
    echo "  --with-job      包含 Job 服务"
    echo "  --with-consumer 包含 Consumer 服务"
    echo "  -a, --all       包含所有可选服务"
    echo "  -i, --interactive 交互式模式"
    echo "  -h, --help      显示帮助"
    echo ""
    echo "示例:"
    echo "  ./scripts/init-project.sh -i                     # 交互式模式"
    echo "  ./scripts/init-project.sh -n myapp -m github.com/org/myapp"
    echo "  ./scripts/init-project.sh -n myapp -a            # 包含所有服务"
    echo "  ./scripts/init-project.sh -n myapp --with-rpc    # 只包含 RPC"
}

# 解析参数
while [[ $# -gt 0 ]]; do
    case $1 in
        -n|--name) PROJECT_NAME="$2"; shift 2 ;;
        -m|--module) MODULE_PATH="$2"; shift 2 ;;
        --with-rpc) WITH_RPC=true; shift ;;
        --with-job) WITH_JOB=true; shift ;;
        --with-consumer) WITH_CONSUMER=true; shift ;;
        -a|--all) WITH_RPC=true; WITH_JOB=true; WITH_CONSUMER=true; shift ;;
        -i|--interactive) INTERACTIVE=true; shift ;;
        -h|--help) show_help; exit 0 ;;
        *) echo -e "${RED}未知参数: $1${NC}"; show_help; exit 1 ;;
    esac
done

# 交互式模式
if [ "$INTERACTIVE" = true ]; then
    echo -e "${GREEN}🚀 IDRM 项目初始化向导${NC}"
    echo ""
    
    read -p "项目名称 [my-project]: " input
    PROJECT_NAME=${input:-$PROJECT_NAME}
    
    read -p "Go 模块路径 [github.com/example/$PROJECT_NAME]: " input
    MODULE_PATH=${input:-"github.com/example/$PROJECT_NAME"}
    
    echo ""
    echo "选择要包含的服务组件:"
    
    read -p "包含 RPC 服务? (y/N): " input
    [[ "$input" =~ ^[Yy]$ ]] && WITH_RPC=true
    
    read -p "包含 Job 服务? (y/N): " input
    [[ "$input" =~ ^[Yy]$ ]] && WITH_JOB=true
    
    read -p "包含 Consumer 服务? (y/N): " input
    [[ "$input" =~ ^[Yy]$ ]] && WITH_CONSUMER=true
fi

# 设置默认模块路径
[ -z "$MODULE_PATH" ] && MODULE_PATH="github.com/example/$PROJECT_NAME"

echo ""
echo -e "${GREEN}🚀 初始化项目: $PROJECT_NAME${NC}"
echo "   模块路径: $MODULE_PATH"
echo ""

# 删除未选择的可选服务
if [ "$WITH_RPC" = false ]; then
    echo -e "${YELLOW}   跳过 RPC 服务${NC}"
    rm -rf rpc/
fi
if [ "$WITH_JOB" = false ]; then
    echo -e "${YELLOW}   跳过 Job 服务${NC}"
    rm -rf job/
fi
if [ "$WITH_CONSUMER" = false ]; then
    echo -e "${YELLOW}   跳过 Consumer 服务${NC}"
    rm -rf consumer/
fi

# 替换占位符
echo "   替换占位符..."
find . -type f \( -name "*.go" -o -name "*.md" -o -name "*.yaml" -o -name "*.api" -o -name "*.template" \) \
    -exec sed -i '' "s/{{PROJECT_NAME}}/$PROJECT_NAME/g" {} \; 2>/dev/null || true
find . -type f \( -name "*.go" -o -name "*.md" -o -name "*.yaml" \) \
    -exec sed -i '' "s|{{MODULE_PATH}}|$MODULE_PATH|g" {} \; 2>/dev/null || true

# 初始化 go.mod
if [ -f "go.mod.template" ]; then
    echo "   初始化 go.mod..."
    mv go.mod.template go.mod
    sed -i '' "s|{{MODULE_PATH}}|$MODULE_PATH|g" go.mod 2>/dev/null || true
fi

# 重命名配置文件
if [ -f "api/etc/api.yaml.template" ]; then
    mv api/etc/api.yaml.template api/etc/api.yaml
fi

# 创建 .gitkeep 文件
touch migrations/.gitkeep 2>/dev/null || true
touch specs/features/.gitkeep 2>/dev/null || true

# 删除初始化脚本自身
rm -f scripts/init-project.sh

echo ""
echo -e "${GREEN}✅ 项目 $PROJECT_NAME 初始化完成！${NC}"
echo ""
echo -e "📋 ${GREEN}已包含的服务:${NC}"
echo "   ✓ API 服务"
[ "$WITH_RPC" = true ] && echo "   ✓ RPC 服务"
[ "$WITH_JOB" = true ] && echo "   ✓ Job 服务"
[ "$WITH_CONSUMER" = true ] && echo "   ✓ Consumer 服务"
echo ""
echo -e "🔧 ${GREEN}下一步:${NC}"
echo "   1. 配置 api/etc/api.yaml (数据库连接等)"
echo "   2. go mod tidy"
echo "   3. go run api/api.go"
echo ""
echo -e "📖 ${GREEN}文档:${NC}"
echo "   - 阅读 CLAUDE.md 了解开发规范"
echo "   - 阅读 sdd_doc/spec/ 了解项目规范"
