//go:build integration
// +build integration

package rbac

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun"

	"github.com/determined-ai/determined/master/internal/db"
	"github.com/determined-ai/determined/master/internal/usergroup"
	"github.com/determined-ai/determined/master/pkg/model"
	"github.com/determined-ai/determined/proto/pkg/rbacv1"
)

const (
	pathToMigrations  = "file://../../static/migrations"
	testWorkspaceName = "test workspace"
)

var (
	testGroupStatic = usergroup.Group{
		ID:   10001,
		Name: "testGroupStatic",
	}
	testGroupOwnedByUser = usergroup.Group{
		ID:      9999,
		Name:    "testGroupUser",
		OwnerID: 1217651234,
	}
	testUser = model.User{
		ID:       1217651234,
		Username: fmt.Sprintf("IntegrationTest%d", 1217651234),
		Admin:    false,
		Active:   false,
	}

	testRole = Role{
		ID:   10002,
		Name: "test role 1",
	}
	testRole2 = Role{
		ID:   10003,
		Name: "test role 2",
	}
	testRole3 = Role{
		ID:   10004,
		Name: "test role 3",
	}
	testRole4 = Role{
		ID:   10005,
		Name: "test role 4",
	}
	testRoles = []Role{testRole, testRole2, testRole3, testRole4}

	testPermission = Permission{
		ID:     10006,
		Name:   "test permission 1",
		Global: false,
	}
	testPermission2 = Permission{
		ID:     10007,
		Name:   "test permission 2",
		Global: false,
	}
	testPermission3 = Permission{
		ID:     10008,
		Name:   "test permission 3",
		Global: false,
	}
	globalTestPermission = Permission{
		ID:     10009,
		Name:   "test permission global",
		Global: true,
	}
	testPermissions = []Permission{testPermission, testPermission2, testPermission3, globalTestPermission}

	testWorkspace = Workspace{
		ID:   10011,
		Name: "test workspace",
	}

	testRoleAssignment = RoleAssignmentScope{
		ID: 10010,
		WorkspaceID: sql.NullInt32{
			Int32: int32(testWorkspace.ID),
			Valid: true,
		},
	}
)

type Workspace struct {
	bun.BaseModel `bun:"table:workspaces"`

	ID       int    `bun:"id,notnull"`
	Name     string `bun:"name,notnull"`
	Archived bool   `bun:"archived,notnull"`
}

func TestRbac(t *testing.T) {
	ctx := context.Background()
	pgDB := db.MustResolveTestPostgres(t)
	db.MustMigrateTestPostgres(t, pgDB, pathToMigrations)

	t.Cleanup(func() { cleanUp(ctx, t, pgDB) })
	setUp(ctx, t, pgDB)

	rbacRole := &rbacv1.Role{
		RoleId: int32(testRole.ID),
		Name:   testRole.Name,
		Permissions: []*rbacv1.Permission{
			{
				Id:       int32(testPermission.ID),
				Name:     testPermission.Name,
				IsGlobal: testPermission.Global,
			},
		},
	}
	rbacRole2 := &rbacv1.Role{
		RoleId: int32(testRole2.ID),
		Name:   testRole2.Name,
		Permissions: []*rbacv1.Permission{
			{
				Id:       int32(testPermission.ID),
				Name:     testPermission.Name,
				IsGlobal: testPermission.Global,
			},
		},
	}
	rbacRole3 := &rbacv1.Role{
		RoleId: int32(testRole3.ID),
		Name:   testRole3.Name,
		Permissions: []*rbacv1.Permission{
			{
				Id:       int32(testPermission.ID),
				Name:     testPermission.Name,
				IsGlobal: testPermission.Global,
			},
		},
	}

	workspaceId := int32(testRoleAssignment.WorkspaceID.Int32)
	userRoleAssignment := rbacv1.UserRoleAssignment{
		UserId: int32(testUser.ID),
		RoleAssignment: &rbacv1.RoleAssignment{
			Role:             rbacRole,
			ScopeWorkspaceId: &workspaceId,
		},
	}

	groupRoleAssignment := rbacv1.GroupRoleAssignment{
		GroupId: int32(testGroupStatic.ID),
		RoleAssignment: &rbacv1.RoleAssignment{
			Role:             rbacRole,
			ScopeWorkspaceId: &workspaceId,
		},
	}
	assignmentScope := RoleAssignmentScope{}
	assignment := RoleAssignment{}

	t.Run("test user role assignment", func(t *testing.T) {
		// TODO: populate the permission assignments table in the future
		err := AddRoleAssignments(
			ctx, []*rbacv1.GroupRoleAssignment{}, []*rbacv1.UserRoleAssignment{&userRoleAssignment})
		require.NoError(t, err, "error adding role assigment")

		err = db.Bun().NewSelect().Model(&assignmentScope).Where(
			"scope_workspace_id=?", testRoleAssignment.WorkspaceID.Int32).Scan(ctx)
		require.NoError(t, err, "error getting created assignment scope")

		err = db.Bun().NewSelect().Model(&assignment).Where("group_id=?", testGroupOwnedByUser.ID).Scan(ctx)
		require.NoError(t, err, "error getting created role assignment")
		require.Equal(t, testGroupOwnedByUser.ID, assignment.GroupID, "incorrect group ID was assigned")
		require.Equal(t, testRole.ID, assignment.RoleID, "incorrect role ID was assigned")
		require.Equal(t, assignmentScope.ID, assignment.ScopeID, "incorrect scope ID was assigned")
	})

	t.Run("test delete user role assignment", func(t *testing.T) {
		err := RemoveRoleAssignments(
			ctx, []*rbacv1.GroupRoleAssignment{}, []*rbacv1.UserRoleAssignment{&userRoleAssignment})
		require.NoError(t, err, "error removing role assignment")

		err = db.Bun().NewSelect().Model(&assignmentScope).Where(
			"scope_workspace_id=?", testRoleAssignment.WorkspaceID.Int32).Scan(ctx)
		require.NoError(t, err, "assignment scope should still exist after removal")

		err = db.Bun().NewSelect().Model(&assignment).Where("group_id=?", testGroupOwnedByUser.ID).Scan(ctx)
		require.Errorf(t, err, "assignment should not exist after removal")
		require.True(t, errors.Is(db.MatchSentinelError(err), db.ErrNotFound), "incorrect error returned")
	})

	t.Run("test group role assignment", func(t *testing.T) {
		// TODO: populate the permission assignments table in the future
		err := AddRoleAssignments(
			ctx, []*rbacv1.GroupRoleAssignment{&groupRoleAssignment}, []*rbacv1.UserRoleAssignment{})
		require.NoError(t, err, "error adding role assigment")

		err = db.Bun().NewSelect().Model(&assignmentScope).Where(
			"scope_workspace_id=?", testRoleAssignment.WorkspaceID.Int32).Scan(ctx)
		require.NoError(t, err, "error getting created assignment scope")

		err = db.Bun().NewSelect().Model(&assignment).Where("group_id=?", testGroupStatic.ID).Scan(ctx)
		require.NoError(t, err, "error getting created role assignment")
		require.Equal(t, testGroupStatic.ID, assignment.GroupID, "incorrect group ID was assigned")
		require.Equal(t, testRole.ID, assignment.RoleID, "incorrect role ID was assigned")
		require.Equal(t, assignmentScope.ID, assignment.ScopeID, "incorrect scope ID was assigned")
	})

	t.Run("test delete group role assignment", func(t *testing.T) {
		err := RemoveRoleAssignments(
			ctx, []*rbacv1.GroupRoleAssignment{&groupRoleAssignment}, []*rbacv1.UserRoleAssignment{})
		require.NoError(t, err, "error removing role assignment")

		err = db.Bun().NewSelect().Model(&assignmentScope).Where(
			"scope_workspace_id=?", testRoleAssignment.WorkspaceID.Int32).Scan(ctx)
		require.NoError(t, err, "assignment scope should still exist after removal")

		err = db.Bun().NewSelect().Model(&assignment).Where("group_id=?", testGroupStatic.ID).Scan(ctx)
		require.Errorf(t, err, "assignment should not exist after removal")
		require.True(t, errors.Is(db.MatchSentinelError(err), db.ErrNotFound), "incorrect error returned")
	})

	t.Run("test add role twice", func(t *testing.T) {
		err := AddRoleAssignments(
			ctx, []*rbacv1.GroupRoleAssignment{&groupRoleAssignment}, []*rbacv1.UserRoleAssignment{})
		require.NoError(t, err, "error adding role assigment")
		err = AddRoleAssignments(
			ctx, []*rbacv1.GroupRoleAssignment{&groupRoleAssignment}, []*rbacv1.UserRoleAssignment{})
		require.Error(t, err, "adding the same role assignment should error")
		require.True(t, errors.Is(err, db.ErrDuplicateRecord), "error should be a duplicate record error")
	})

	t.Run("test insert multiple scopes", func(t *testing.T) {
		nilAssignment := &rbacv1.RoleAssignment{ScopeWorkspaceId: nil}
		_, err := getOrCreateRoleAssignmentScopeTx(ctx, nil, nilAssignment)
		require.NoError(t, err, "error with inserting a nil ")

		_, err = getOrCreateRoleAssignmentScopeTx(ctx, nil, nilAssignment)
		require.NoError(t, err, "inserting the same role assignment scope should not fail")

		rows, err := db.Bun().NewSelect().Table("role_assignment_scopes").Where("scope_workspace_id IS NULL").Count(ctx)
		require.Equal(t, 1, rows, "there should only have been one null scope created")
	})

	t.Run("test get all roles with pagination", func(t *testing.T) {
		permissionsToAdd := []map[string]interface{}{
			{
				"permission_id": globalTestPermission.ID,
				"role_id":       testRole4.ID,
			},
			{
				"permission_id": testPermission.ID,
				"role_id":       testRole3.ID,
			},
		}

		for _, p := range permissionsToAdd {
			_, err := db.Bun().NewInsert().Model(&p).TableExpr("permission_assignments").Exec(ctx)
			require.NoError(t, err, "failure inserting permission assignments in local setup")
		}

		roles, _, err := GetAllRoles(ctx, false, 0, 10)
		require.NoError(t, err, "error getting all roles")
		require.Equal(t, 4, len(roles), "incorrect number of roles retrieved")
		require.True(t, compareRoles(testRole, roles[0]),
			"test role 1 is not equivalent to the retrieved role")
		require.True(t, compareRoles(testRole2, roles[1]),
			"test role 2 is not equivalent to the retrieved role")
		require.True(t, compareRoles(testRole3, roles[2]),
			"test role 3 is not equivalent to the retrieved role")
		require.True(t, compareRoles(testRole4, roles[3]),
			"test role 4 is not equivalent to the retrieved role")

		roles, _, err = GetAllRoles(ctx, true, 0, 10)
		require.NoError(t, err, "error getting non-global roles")
		require.Equal(t, 3, len(roles), "incorrect number of non-global roles retrieved")
		require.True(t, compareRoles(testRole, roles[0]), "test role 1 is not equivalent to the retrieved role")
		require.True(t, compareRoles(testRole2, roles[1]), "test role 2 is not equivalent to the retrieved role")
		require.True(t, compareRoles(testRole3, roles[2]), "test role 3 is not equivalent to the retrieved role")

		roles, _, err = GetAllRoles(ctx, false, 0, 2)
		require.NoError(t, err, "error getting roles with limit")
		require.Equal(t, 2, len(roles), "incorrect number of non-global roles retrieved")
		require.True(t, compareRoles(testRole, roles[0]),
			"test role 1 is not equivalent to the retrieved role")
		require.True(t, compareRoles(testRole2, roles[1]),
			"test role 2 is not equivalent to the retrieved role")

		roles, _, err = GetAllRoles(ctx, false, 2, 10)
		require.NoError(t, err, "error getting roles with offset")
		require.Equal(t, 2, len(roles), "incorrect number of non-global roles retrieved")
		require.True(t, compareRoles(testRole3, roles[0]),
			"test role 3 is not equivalent to the retrieved role")
		require.True(t, compareRoles(testRole4, roles[1]),
			"test role 4 is not equivalent to the retrieved role")
	})

	t.Run("test getting roles by id", func(t *testing.T) {
		rolesWithAssignment, err := GetRolesByIDs(ctx, int32(testRole2.ID), int32(testRole4.ID))
		require.NoError(t, err, "error getting roles 2 and 4 by ID")
		require.Equal(t, testRole2.ID, int(rolesWithAssignment[0].Role.RoleId),
			"test role 2 is not equivalent to the retrieved role")
		require.Equal(t, testRole2.Name, rolesWithAssignment[0].Role.Name,
			"test role 2 is not equivalent to the retrieved role")
		require.Equal(t, testRole4.ID, int(rolesWithAssignment[1].Role.RoleId),
			"test role 4 is not equivalent to the retrieved role")
		require.Equal(t, testRole4.Name, rolesWithAssignment[1].Role.Name,
			"test role 4 is not equivalent to the retrieved role")
	})

	t.Run("test getting roles assigned to group", func(t *testing.T) {
		groupRoleAssignments := []*rbacv1.GroupRoleAssignment{
			{
				GroupId: int32(testGroupStatic.ID),
				RoleAssignment: &rbacv1.RoleAssignment{
					Role:             rbacRole2,
					ScopeWorkspaceId: &workspaceId,
				},
			},
			{
				GroupId: int32(testGroupStatic.ID),
				RoleAssignment: &rbacv1.RoleAssignment{
					Role:             rbacRole3,
					ScopeWorkspaceId: &workspaceId,
				},
			},
		}

		err := AddRoleAssignments(
			ctx, groupRoleAssignments, []*rbacv1.UserRoleAssignment{})
		require.NoError(t, err, "error adding role assigments")

		roles, err := GetRolesAssignedToGroupsTx(ctx, nil, int32(testGroupStatic.ID))
		require.NoError(t, err, "error getting roles assigned to group")
		require.Len(t, roles, 3, "incorrect number of roles retrieved")
		require.True(t, compareRoles(testRole, roles[0]),
			"testRole is not the first role retrieved by group id")
		require.True(t, compareRoles(testRole2, roles[1]),
			"testRole2 is not the second role retrieved by group id")
		require.True(t, compareRoles(testRole3, roles[2]),
			"testRole3 is not the first role retrieved by group id")

		err = RemoveRoleAssignments(ctx, groupRoleAssignments, nil)
		require.NoError(t, err, "error removing assignments from group")
		roles, err = GetRolesAssignedToGroupsTx(ctx, nil, int32(testGroupStatic.ID))
		require.Equal(t, 1, len(roles), "incorrect number of roles retrieved")
	})

	t.Run("test UserPermissionsForScope", func(t *testing.T) {
		groupRoleAssignments := []*rbacv1.GroupRoleAssignment{
			{
				GroupId: int32(testGroupStatic.ID),
				RoleAssignment: &rbacv1.RoleAssignment{
					Role: rbacRole,
				},
			},
			{
				GroupId: int32(testGroupOwnedByUser.ID),
				RoleAssignment: &rbacv1.RoleAssignment{
					Role:             rbacRole2,
					ScopeWorkspaceId: &workspaceId,
				},
			},
		}

		permissionAssignments := []PermissionAssignment{
			{
				PermissionID: globalTestPermission.ID,
				RoleID:       testRole.ID,
			},
			{
				PermissionID: testPermission.ID,
				RoleID:       testRole.ID,
			},
			{
				PermissionID: testPermission2.ID,
				RoleID:       testRole2.ID,
			},
			{
				PermissionID: testPermission3.ID,
				RoleID:       testRole2.ID,
			},
		}

		t.Cleanup(func() {
			// clean out role assignments
			err := RemoveRoleAssignments(ctx, groupRoleAssignments, nil)
			require.NoError(t, err, "error removing group role assignments during cleanup")

			// clean out permission assignments
			_, err = db.Bun().NewDelete().Model(&permissionAssignments).WherePK().Exec(ctx)
			require.NoError(t, err, "error removing permission assignments during cleanup")
		})

		err := AddRoleAssignments(ctx, groupRoleAssignments, nil)
		require.NoError(t, err, "error adding role assigments")

		_, err = db.Bun().NewInsert().Model(&permissionAssignments).Exec(ctx)
		require.NoError(t, err, "error adding permission assignments during setup")

		// Test for non-existent users
		permissions, err := UserPermissionsForScope(ctx, -9999, 0)
		require.NoError(t, err,
			"unexpected error from UserPermissionsForScope when non-existent user")
		require.Empty(t, permissions, "Expected empty permissions for non-existent user")

		// Test for scope-assigned role
		permissions, err = UserPermissionsForScope(ctx, testUser.ID, testWorkspace.ID)
		require.Len(t, permissions, 4, "Expected four permissions from %v", permissions)
		require.True(t, permissionsContainsAll(permissions,
			globalTestPermission.ID, testPermission.ID, testPermission2.ID, testPermission3.ID),
			"failed to find expected permissions for scope-assigned role in %v", permissions)

		// Test for globally assigned role
		permissions, err = UserPermissionsForScope(ctx, testUser.ID, 0)
		require.Len(t, permissions, 2, "Expected two permissions from %v", permissions)
		require.True(t, permissionsContainsAll(permissions,
			globalTestPermission.ID, testPermission.ID), "failed to find expected permissions in %v", permissions)
		require.False(t, permissionsContainsAll(permissions, testPermission2.ID),
			"Unexpectedly found permission %v for user in %v", testPermission2.ID, permissions)
		require.False(t, permissionsContainsAll(permissions, testPermission3.ID),
			"Unexpectedly found permission %v for user in %v", testPermission3.ID, permissions)
	})
}

func setUp(ctx context.Context, t *testing.T, pgDB *db.PgDB) {
	_, err := pgDB.AddUser(&testUser, nil)
	require.NoError(t, err, "failure creating user in setup")

	_, _, err = usergroup.AddGroupWithMembers(ctx, testGroupStatic, testUser.ID)
	require.NoError(t, err, "failure creating static test group")

	testGroupOwnedByUser.OwnerID = testUser.ID
	_, _, err = usergroup.AddGroupWithMembers(ctx, testGroupOwnedByUser, testUser.ID)
	require.NoError(t, err, "failure creating test user group")

	_, err = db.Bun().NewInsert().Model(&testPermissions).Exec(ctx)
	require.NoError(t, err, "failure creating permission in setup")

	_, err = db.Bun().NewInsert().Model(&testRoles).Exec(ctx)
	require.NoError(t, err, "failure creating role in setup")

	workspace := map[string]interface{}{
		"name": testWorkspace.Name,
		"id":   testWorkspace.ID,
	}
	_, err = db.Bun().NewInsert().Model(&workspace).TableExpr("workspaces").Exec(ctx)
	require.NoError(t, err, "failure creating workspace in setup")
}

func cleanUp(ctx context.Context, t *testing.T, pgDB *db.PgDB) {
	_, err := db.Bun().NewDelete().Table("workspaces").Where(
		"name=?", "test workspace").Exec(ctx)
	if err != nil {
		t.Logf("Error cleaning up workspace")
	}

	_, err = db.Bun().NewDelete().Model(&testPermissions).WherePK().Exec(ctx)
	if err != nil {
		t.Logf("Error cleaning up permissions")
	}

	_, err = db.Bun().NewDelete().Table("roles").Where("id IN (?)",
		bun.In([]int32{
			int32(testRole.ID), int32(testRole2.ID),
			int32(testRole3.ID), int32(testRole4.ID),
		})).Exec(ctx)
	if err != nil {
		t.Logf("Error cleaning up role")
	}

	err = usergroup.RemoveUsersFromGroupTx(ctx, nil, testGroupStatic.ID, testUser.ID)
	if err != nil {
		t.Logf("Error cleaning up group membership on (%v, %v): %v", testGroupStatic.ID, testUser.ID, err)
	}

	_, err = db.Bun().NewDelete().Table("role_assignments").Where(
		"group_id=?", testGroupStatic.ID).Exec(ctx)
	if err != nil {
		t.Log("Error cleaning up static group from role assignment")
	}

	err = usergroup.DeleteGroup(ctx, testGroupStatic.ID)
	if err != nil {
		t.Logf("Error cleaning up static group: %v", err)
	}

	_, err = db.Bun().NewDelete().Table("users").Where("id = ?", testUser.ID).Exec(ctx)
	if err != nil {
		t.Logf("Error cleaning up user: %v\n", err)
	}
}

func compareRoles(expected, actual Role) bool {
	switch {
	case expected.ID != actual.ID:
		return false
	case expected.Name != actual.Name:
		return false
	case !expected.Created.Equal(actual.Created):
		return false
	}
	return true
}

func permissionsContainsAll(permissions []Permission, ids ...int) bool {
	foundIDs := make(map[int]bool)

	for _, id := range ids {
		for _, p := range permissions {
			if p.ID == id {
				foundIDs[id] = true
			}
		}
	}

	for _, id := range ids {
		if !foundIDs[id] {
			return false
		}
	}

	return true
}
