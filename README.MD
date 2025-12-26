# SimpleHRSystem

[![Go Version](https://img.shields.io/badge/Go-1.2x-blue.svg)](https://golang.org/)
[![Framework](https://img.shields.io/badge/Framework-GoFrame-brightgreen.svg)](https://goframe.org/)
[![Database](https://img.shields.io/badge/DB-MySQL-orange.svg)](https://www.mysql.com/)
[![Cache](https://img.shields.io/badge/Cache-Redis-red.svg)](https://redis.io/)

SimpleHRSystem 是一款基于 GoFrame 框架开发的轻量级人力资源管理系统，专为 Go 语言新手打造的练手项目，覆盖后端开发常见技术栈（GoFrame+Redis+MySQL），实现HR管理核心基础功能，帮助新手理解企业级Go项目的开发流程、分层设计和中间件使用。

## 🌟 项目特点
- 基于 GoFrame 最新版开发，遵循 GF 框架最佳实践
- 分层架构设计（API层、Service层、Dao层），代码结构清晰
- 集成 Redis 实现热点数据缓存（如部门列表、员工基础信息）
- MySQL 数据持久化，使用 GF 自带的 ORM 简化数据库操作
- 完整的接口示例（RESTful API），包含参数校验、异常处理
- 适合 Go 1.2x 版本，兼容最新语法特性

## 📋 核心功能
### 1. 员工管理
- 员工信息的增删改查（CRUD）
- 员工状态（在职/离职/休假）管理
- 员工基本信息导入/导出（简易版）

### 2. 部门管理
- 部门层级结构维护
- 部门人员关联查询
- 部门信息缓存（Redis）

### 3. 考勤管理
- 考勤记录新增/查询
- 月度考勤统计
- 考勤异常标记

### 4. 基础功能
- 接口统一返回格式
- 全局异常处理
- 日志记录（GF 日志组件）
- 参数自动校验

## 🛠 技术栈
| 组件         | 版本/说明                |
|--------------|--------------------------|
| Go           | 1.20+（推荐1.21/1.22）|
| 框架         | GoFrame（gf）v2          |
| 数据库       | MySQL 8.0+               |
| 缓存         | Redis 6.0+               |
| ORM          | GF 自带 gdb              |
| 依赖管理     | Go Mod                   |

## 🚀 快速开始
### 1. 环境准备
- 安装 Go 1.20+（确认 `go version` 输出符合要求）
- 安装 MySQL 8.0+，创建数据库 `simple_hr_system`
- 安装 Redis 6.0+，确保服务正常运行（默认端口6379）

### 2. 克隆项目
```bash
git clone https://github.com/yxb123456cy/SimpleHRSystem.git
cd SimpleHRSystem