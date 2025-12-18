<!--
  角色管理页面
  角色列表展示和管理
-->
<template>
  <div class="role-container">
    <el-card>
      <div class="card-header">
        <span>角色列表</span>
        <el-button type="primary" @click="handleAdd">
          <el-icon><Plus /></el-icon>
          新增角色
        </el-button>
      </div>
      
      <el-table :data="roleList" border stripe>
        <el-table-column prop="roleId" label="角色ID" width="80" />
        <el-table-column prop="roleName" label="角色名称" />
        <el-table-column prop="roleKey" label="权限字符" />
        <el-table-column prop="roleSort" label="排序" width="80" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === 0 ? 'success' : 'danger'">
              {{ row.status === 0 ? '正常' : '停用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="createTime" label="创建时间" width="180" />
        <el-table-column label="操作" fixed="right" width="200">
          <template #default="{ row }">
            <el-button type="primary" link>编辑</el-button>
            <el-button type="danger" link>删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const roleList = ref([])

const handleAdd = () => {
  console.log('新增角色')
}
</script>

<style lang="scss" scoped>
.role-container {
  .card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 20px;
  }
}
</style>
