// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/github/user/events": {
            "get": {
                "description": "获取开发者活动数据",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Github"
                ],
                "summary": "获取开发者活动数据",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Github用户名",
                        "name": "username",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github.EventListResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.OperationResp"
                        }
                    }
                }
            }
        },
        "/github/user/info": {
            "get": {
                "description": "获取开发者基本信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Github"
                ],
                "summary": "获取开发者基本信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Github用户名",
                        "name": "username",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github.UserResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.OperationResp"
                        }
                    }
                }
            }
        },
        "/github/user/nation": {
            "get": {
                "description": "获取开发者国籍",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Github"
                ],
                "summary": "获取开发者国籍",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Github用户名",
                        "name": "username",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.DevNationalityResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.OperationResp"
                        }
                    }
                }
            }
        },
        "/github/user/repos": {
            "get": {
                "description": "获取仓库和语言偏好",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Github"
                ],
                "summary": "获取仓库和语言偏好",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Github用户名",
                        "name": "username",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github.RepoListResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.OperationResp"
                        }
                    }
                }
            }
        },
        "/score/get": {
            "get": {
                "description": "获取开发者评分",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Score"
                ],
                "summary": "获取开发者评分",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Github用户名",
                        "name": "username",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.DevScoreResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.OperationResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github.Actor": {
            "type": "object",
            "properties": {
                "avatar_url": {
                    "type": "string"
                },
                "display_login": {
                    "type": "string"
                },
                "gravatar_id": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "login": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "github.Commit": {
            "type": "object",
            "properties": {
                "author": {
                    "$ref": "#/definitions/github.User"
                },
                "comments_url": {
                    "type": "string"
                },
                "commit": {
                    "$ref": "#/definitions/github.CommitDetails"
                },
                "committer": {
                    "$ref": "#/definitions/github.User"
                },
                "distinct": {
                    "type": "boolean"
                },
                "html_url": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "node_id": {
                    "type": "string"
                },
                "parents": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github.Parent"
                    }
                },
                "sha": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "github.CommitDetails": {
            "type": "object",
            "properties": {
                "author": {
                    "$ref": "#/definitions/github.User"
                },
                "comment_count": {
                    "type": "integer"
                },
                "committer": {
                    "$ref": "#/definitions/github.User"
                },
                "message": {
                    "type": "string"
                },
                "tree": {
                    "$ref": "#/definitions/github.Tree"
                },
                "url": {
                    "type": "string"
                },
                "verification": {
                    "$ref": "#/definitions/github.Verification"
                }
            }
        },
        "github.Event": {
            "type": "object",
            "properties": {
                "actor": {
                    "$ref": "#/definitions/github.Actor"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "payload": {
                    "$ref": "#/definitions/github.Payload"
                },
                "public": {
                    "type": "boolean"
                },
                "repo": {
                    "$ref": "#/definitions/github.Repository"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "github.EventListResp": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github.Event"
                    }
                }
            }
        },
        "github.License": {
            "type": "object",
            "properties": {
                "html_url": {
                    "type": "string"
                },
                "key": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "node_id": {
                    "type": "string"
                },
                "spdx_id": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "github.Parent": {
            "type": "object",
            "properties": {
                "sha": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "github.Payload": {
            "type": "object",
            "properties": {
                "action": {
                    "type": "string"
                },
                "before": {
                    "type": "string"
                },
                "commits": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github.Commit"
                    }
                },
                "distinct_size": {
                    "type": "integer"
                },
                "head": {
                    "type": "string"
                },
                "push_id": {
                    "type": "integer"
                },
                "ref": {
                    "type": "string"
                },
                "size": {
                    "type": "integer"
                }
            }
        },
        "github.Permissions": {
            "type": "object",
            "properties": {
                "admin": {
                    "type": "boolean"
                },
                "pull": {
                    "type": "boolean"
                },
                "push": {
                    "type": "boolean"
                }
            }
        },
        "github.Repo": {
            "type": "object",
            "properties": {
                "allow_auto_merge": {
                    "type": "boolean"
                },
                "allow_forking": {
                    "type": "boolean"
                },
                "allow_merge_commit": {
                    "type": "boolean"
                },
                "allow_rebase_merge": {
                    "type": "boolean"
                },
                "allow_squash_merge": {
                    "type": "boolean"
                },
                "archive_url": {
                    "type": "string"
                },
                "archived": {
                    "type": "boolean"
                },
                "assignees_url": {
                    "type": "string"
                },
                "blobs_url": {
                    "type": "string"
                },
                "branches_url": {
                    "type": "string"
                },
                "clone_url": {
                    "type": "string"
                },
                "collaborators_url": {
                    "type": "string"
                },
                "comments_url": {
                    "type": "string"
                },
                "commits_url": {
                    "type": "string"
                },
                "compare_url": {
                    "type": "string"
                },
                "contents_url": {
                    "type": "string"
                },
                "contributors_url": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "default_branch": {
                    "type": "string"
                },
                "delete_branch_on_merge": {
                    "type": "boolean"
                },
                "deployments_url": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "disabled": {
                    "type": "boolean"
                },
                "downloads_url": {
                    "type": "string"
                },
                "events_url": {
                    "type": "string"
                },
                "fork": {
                    "type": "boolean"
                },
                "forks": {
                    "type": "integer"
                },
                "forks_count": {
                    "type": "integer"
                },
                "forks_url": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "git_commits_url": {
                    "type": "string"
                },
                "git_refs_url": {
                    "type": "string"
                },
                "git_tags_url": {
                    "type": "string"
                },
                "git_url": {
                    "type": "string"
                },
                "has_discussions": {
                    "type": "boolean"
                },
                "has_downloads": {
                    "type": "boolean"
                },
                "has_issues": {
                    "type": "boolean"
                },
                "has_pages": {
                    "type": "boolean"
                },
                "has_projects": {
                    "type": "boolean"
                },
                "has_wiki": {
                    "type": "boolean"
                },
                "homepage": {
                    "type": "string"
                },
                "html_url": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_template": {
                    "type": "boolean"
                },
                "issue_comment_url": {
                    "type": "string"
                },
                "issue_events_url": {
                    "type": "string"
                },
                "issues_url": {
                    "type": "string"
                },
                "keys_url": {
                    "type": "string"
                },
                "labels_url": {
                    "type": "string"
                },
                "languages_url": {
                    "type": "string"
                },
                "license": {
                    "$ref": "#/definitions/github.License"
                },
                "merges_url": {
                    "type": "string"
                },
                "milestones_url": {
                    "type": "string"
                },
                "mirror_url": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "network_count": {
                    "type": "integer"
                },
                "node_id": {
                    "type": "string"
                },
                "notifications_url": {
                    "type": "string"
                },
                "open_issues": {
                    "type": "integer"
                },
                "open_issues_count": {
                    "type": "integer"
                },
                "organization": {
                    "$ref": "#/definitions/github.User"
                },
                "owner": {
                    "$ref": "#/definitions/github.User"
                },
                "parent": {
                    "$ref": "#/definitions/github.Repo"
                },
                "permissions": {
                    "$ref": "#/definitions/github.Permissions"
                },
                "private": {
                    "type": "boolean"
                },
                "pulls_url": {
                    "type": "string"
                },
                "pushed_at": {
                    "type": "string"
                },
                "releases_url": {
                    "type": "string"
                },
                "size": {
                    "type": "integer"
                },
                "source": {
                    "$ref": "#/definitions/github.Repo"
                },
                "ssh_url": {
                    "type": "string"
                },
                "stargazers_count": {
                    "type": "integer"
                },
                "stargazers_url": {
                    "type": "string"
                },
                "statuses_url": {
                    "type": "string"
                },
                "subscribers_count": {
                    "type": "integer"
                },
                "subscribers_url": {
                    "type": "string"
                },
                "subscription_url": {
                    "type": "string"
                },
                "tags_url": {
                    "type": "string"
                },
                "teams_url": {
                    "type": "string"
                },
                "temp_clone_token": {
                    "type": "string"
                },
                "template_repository": {
                    "$ref": "#/definitions/github.TemplateRepository"
                },
                "topics": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "trees_url": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                },
                "visibility": {
                    "type": "string"
                },
                "watchers": {
                    "type": "integer"
                },
                "watchers_count": {
                    "type": "integer"
                }
            }
        },
        "github.RepoListResp": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github.Repo"
                    }
                }
            }
        },
        "github.Repository": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "github.TemplateRepository": {
            "type": "object",
            "properties": {
                "allow_auto_merge": {
                    "type": "boolean"
                },
                "allow_merge_commit": {
                    "type": "boolean"
                },
                "allow_rebase_merge": {
                    "type": "boolean"
                },
                "allow_squash_merge": {
                    "type": "boolean"
                },
                "archive_url": {
                    "type": "string"
                },
                "archived": {
                    "type": "boolean"
                },
                "assignees_url": {
                    "type": "string"
                },
                "blobs_url": {
                    "type": "string"
                },
                "branches_url": {
                    "type": "string"
                },
                "clone_url": {
                    "type": "string"
                },
                "collaborators_url": {
                    "type": "string"
                },
                "comments_url": {
                    "type": "string"
                },
                "commits_url": {
                    "type": "string"
                },
                "compare_url": {
                    "type": "string"
                },
                "contents_url": {
                    "type": "string"
                },
                "contributors_url": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "default_branch": {
                    "type": "string"
                },
                "delete_branch_on_merge": {
                    "type": "boolean"
                },
                "deployments_url": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "disabled": {
                    "type": "boolean"
                },
                "downloads_url": {
                    "type": "string"
                },
                "events_url": {
                    "type": "string"
                },
                "fork": {
                    "type": "boolean"
                },
                "forks": {
                    "type": "integer"
                },
                "forks_count": {
                    "type": "integer"
                },
                "forks_url": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "git_commits_url": {
                    "type": "string"
                },
                "git_refs_url": {
                    "type": "string"
                },
                "git_tags_url": {
                    "type": "string"
                },
                "git_url": {
                    "type": "string"
                },
                "has_downloads": {
                    "type": "boolean"
                },
                "has_issues": {
                    "type": "boolean"
                },
                "has_pages": {
                    "type": "boolean"
                },
                "has_projects": {
                    "type": "boolean"
                },
                "has_wiki": {
                    "type": "boolean"
                },
                "homepage": {
                    "type": "string"
                },
                "html_url": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_template": {
                    "type": "boolean"
                },
                "issue_comment_url": {
                    "type": "string"
                },
                "issue_events_url": {
                    "type": "string"
                },
                "issues_url": {
                    "type": "string"
                },
                "keys_url": {
                    "type": "string"
                },
                "labels_url": {
                    "type": "string"
                },
                "language": {},
                "languages_url": {
                    "type": "string"
                },
                "license": {
                    "$ref": "#/definitions/github.License"
                },
                "merges_url": {
                    "type": "string"
                },
                "milestones_url": {
                    "type": "string"
                },
                "mirror_url": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "network_count": {
                    "type": "integer"
                },
                "node_id": {
                    "type": "string"
                },
                "notifications_url": {
                    "type": "string"
                },
                "open_issues": {
                    "type": "integer"
                },
                "open_issues_count": {
                    "type": "integer"
                },
                "owner": {
                    "$ref": "#/definitions/github.User"
                },
                "permissions": {
                    "$ref": "#/definitions/github.Permissions"
                },
                "private": {
                    "type": "boolean"
                },
                "pulls_url": {
                    "type": "string"
                },
                "pushed_at": {
                    "type": "string"
                },
                "releases_url": {
                    "type": "string"
                },
                "size": {
                    "type": "integer"
                },
                "ssh_url": {
                    "type": "string"
                },
                "stargazers_count": {
                    "type": "integer"
                },
                "stargazers_url": {
                    "type": "string"
                },
                "statuses_url": {
                    "type": "string"
                },
                "subscribers_count": {
                    "type": "integer"
                },
                "subscribers_url": {
                    "type": "string"
                },
                "subscription_url": {
                    "type": "string"
                },
                "tags_url": {
                    "type": "string"
                },
                "teams_url": {
                    "type": "string"
                },
                "temp_clone_token": {
                    "type": "string"
                },
                "topics": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "trees_url": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                },
                "visibility": {
                    "type": "string"
                },
                "watchers": {
                    "type": "integer"
                },
                "watchers_count": {
                    "type": "integer"
                }
            }
        },
        "github.Tree": {
            "type": "object",
            "properties": {
                "sha": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "github.User": {
            "type": "object",
            "properties": {
                "avatar_url": {
                    "type": "string"
                },
                "events_url": {
                    "type": "string"
                },
                "followers_url": {
                    "type": "string"
                },
                "following_url": {
                    "type": "string"
                },
                "gists_url": {
                    "type": "string"
                },
                "gravatar_id": {
                    "type": "string"
                },
                "html_url": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "login": {
                    "type": "string"
                },
                "node_id": {
                    "type": "string"
                },
                "organizations_url": {
                    "type": "string"
                },
                "received_events_url": {
                    "type": "string"
                },
                "repos_url": {
                    "type": "string"
                },
                "site_admin": {
                    "type": "boolean"
                },
                "starred_url": {
                    "type": "string"
                },
                "subscriptions_url": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "github.UserInfo": {
            "type": "object",
            "properties": {
                "avatar_url": {
                    "type": "string"
                },
                "bio": {
                    "type": "string"
                },
                "blog": {
                    "type": "string"
                },
                "collaborators": {
                    "type": "integer"
                },
                "company": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "disk_usage": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "events_url": {
                    "type": "string"
                },
                "followers": {
                    "type": "integer"
                },
                "followers_url": {
                    "type": "string"
                },
                "following": {
                    "type": "integer"
                },
                "following_url": {
                    "type": "string"
                },
                "gists_url": {
                    "type": "string"
                },
                "gravatar_id": {
                    "type": "string"
                },
                "hireable": {
                    "type": "boolean"
                },
                "html_url": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "location": {
                    "type": "string"
                },
                "login": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "node_id": {
                    "type": "string"
                },
                "organizations_url": {
                    "type": "string"
                },
                "owned_private_repos": {
                    "type": "integer"
                },
                "private_gists": {
                    "type": "integer"
                },
                "public_gists": {
                    "type": "integer"
                },
                "public_repos": {
                    "type": "integer"
                },
                "received_events_url": {
                    "type": "string"
                },
                "repos_url": {
                    "type": "string"
                },
                "site_admin": {
                    "type": "boolean"
                },
                "starred_url": {
                    "type": "string"
                },
                "subscriptions_url": {
                    "type": "string"
                },
                "total_private_repos": {
                    "type": "integer"
                },
                "two_factor_authentication": {
                    "type": "boolean"
                },
                "type": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "github.UserResp": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "user": {
                    "$ref": "#/definitions/github.UserInfo"
                }
            }
        },
        "github.Verification": {
            "type": "object",
            "properties": {
                "payload": {
                    "type": "string"
                },
                "reason": {
                    "type": "string"
                },
                "signature": {
                    "type": "string"
                },
                "verified": {
                    "type": "boolean"
                }
            }
        },
        "model.DevNationalityResp": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "nation": {
                    "type": "string"
                }
            }
        },
        "model.DevScoreResp": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "score": {
                    "$ref": "#/definitions/model.DeveloperRank"
                }
            }
        },
        "model.DeveloperRank": {
            "type": "object",
            "properties": {
                "score": {
                    "$ref": "#/definitions/model.Score"
                },
                "username": {
                    "description": "开发者用户名",
                    "type": "string"
                }
            }
        },
        "model.OperationResp": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "model.Score": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "代码贡献量评分",
                    "type": "number"
                },
                "influence": {
                    "description": "社区影响力评分",
                    "type": "number"
                },
                "overall": {
                    "description": "综合评分",
                    "type": "number"
                },
                "project": {
                    "description": "项目重要性评分",
                    "type": "number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
