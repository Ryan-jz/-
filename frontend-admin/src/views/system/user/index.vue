<!--
  用户管理页面
  用户列表展示、搜索、新增、编辑、删除功能
-->
<template>
  <div class="user-container">
    <el-card>
      <!-- 搜索栏 -->
      <el-form :inline="true" :model="queryForm" class="search-form">
        <el-form-item label="用户名">
          <el-input v-model="queryForm.userName" placeholder="请输入用户名" clearable />
        </el-form-item>
        <el-form-item label="手机号">
          <el-input v-model="queryForm.phonenumber" placeholder="请输入手机号" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="queryForm.status" placeholder="请选择状态" clearable>
            <el-option label="正常" :value="0" />
            <el-option label="停用" :value="1" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleQuery">
            <el-icon><Search /></el-icon>
            搜索
          </el-button>
          <el-button @click="handleReset">
            <el-icon><Refresh /></el-icon>
            重置
          </el-button>
        </el-form-item>
      </el-form>

      <!-- 操作按钮 -->
      <el-row class="toolbar">
        <el-button type="primary" @click="handleAdd">
          <el-icon><Plus /></el-icon>
          新增用户
        </el-button>
      </el-row>

      <!-- 用户列表 -->
      <el-table :data="userList" v-loading="loading" border stripe>
        <el-table-column prop="userId" label="用户ID" width="80" />
        <el-table-column prop="userName" label="用户名" width="120" />
        <el-table-column prop="nickName" label="昵称" width="120" />
        <el-table-column prop="email" label="邮箱" width="180" />
        <el-table-column prop="phonenumber" label="手机号" width="120" />
        <el-table-column prop="sex" label="性别" width="80">
          <template #default="{ row }">
            <span>{{ row.sex === '0' ? '男' : row.sex === '1' ? '女' : '未知' }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.status === 0 ? 'success' : 'danger'">
              {{ row.status === 0 ? '正常' : '停用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="createTime" label="创建时间" width="180" />
        <el-table-column label="操作" fixed="right" width="200">
          <template #default="{ row }">
            <el-button type="primary" link @click="handleEdit(row)">
              <el-icon><Edit /></el-icon>
              编辑
            </el-button>
            <el-button type="danger" link @click="handleDelete(row)">
              <el-icon><Delete /></el-icon>
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- 分页 -->
      <el-pagination
        v-model:current-page="queryForm.page"
        v-model:page-size="queryForm.pageSize"
        :total="total"
        :page-sizes="[10, 20, 50, 100]"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="handleQuery"
        @current-change="handleQuery"
        class="pagination"
      />
    </el-card>

    <!-- 新增/编辑对话框 -->
    <el-dialog
      v-model="dialogVisible"
      :title="dialogTitle"
      width="600px"
      @close="handleDialogClose"
    >
      <el-form
        ref="userFormRef"
        :model="userForm"
        :rules="userRules"
        label-width="80px"
      >
        <el-form-item label="用户名" prop="userName">
          <el-input v-model="userForm.userName" :disabled="isEdit" />
        </el-form-item>
        <el-form-item label="昵称" prop="nickName">
          <el-input v-model="userForm.nickName" />
        </el-form-item>
        <el-form-item label="密码" prop="password" v-if="!isEdit">
          <el-input v-model="userForm.password" type="password" show-password />
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="userForm.email" />
        </el-form-item>
        <el-form-item label="手机号" prop="phonenumber">
          <el-input v-model="userForm.phonenumber" />
        </el-form-item>
        <el-form-item label="性别" prop="sex">
          <el-radio-group v-model="userForm.sex">
            <el-radio label="0">男</el-radio>
            <el-radio label="1">女</el-radio>
            <el-radio label="2">未知</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="userForm.status">
            <el-radio :label="0">正常</el-radio>
            <el-radio :label="1">停用</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="userForm.remark" type="textarea" :rows="3" />
        </el-form-item>
      </el-form>
      
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleSubmit" :loading="submitLoading">
          确定
        </el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getUserList, createUser, updateUser, deleteUser } from '@/api/user'

// 查询表单
const queryForm = reactive({
  userName: '',
  phonenumber: '',
  status: null,
  page: 1,
  pageSize: 10
})

// 用户列表
const userList = ref([])
const total = ref(0)
const loading = ref(false)

// 对话框
const dialogVisible = ref(false)
const dialogTitle = ref('新增用户')
const isEdit = ref(false)
const submitLoading = ref(false)

// 用户表单
const userFormRef = ref(null)
const userForm = reactive({
  userId: null,
  userName: '',
  nickName: '',
  password: '',
  email: '',
  phonenumber: '',
  sex: '0',
  status: 0,
  deptId: 100,
  remark: ''
})

// 表单验证规则
const userRules = {
  userName: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  nickName: [{ required: true, message: '请输入昵称', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }]
}

/**
 * 获取用户列表
 */
const getList = async () => {
  loading.value = true
  try {
    const res = await getUserList(queryForm)
    userList.value = res.data.list || []
    total.value = res.data.total || 0
  } catch (error) {
    console.error('获取用户列表失败：', error)
  } finally {
    loading.value = false
  }
}

/**
 * 搜索
 */
const handleQuery = () => {
  queryForm.page = 1
  getList()
}

/**
 * 重置
 */
const handleReset = () => {
  queryForm.userName = ''
  queryForm.phonenumber = ''
  queryForm.status = null
  handleQuery()
}

/**
 * 新增用户
 */
const handleAdd = () => {
  isEdit.value = false
  dialogTitle.value = '新增用户'
  dialogVisible.value = true
}

/**
 * 编辑用户
 */
const handleEdit = (row) => {
  isEdit.value = true
  dialogTitle.value = '编辑用户'
  Object.assign(userForm, row)
  dialogVisible.value = true
}

/**
 * 删除用户
 */
const handleDelete = (row) => {
  ElMessageBox.confirm('确定要删除该用户吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      await deleteUser(row.userId)
      ElMessage.success('删除成功')
      getList()
    } catch (error) {
      console.error('删除失败：', error)
    }
  })
}

/**
 * 提交表单
 */
const handleSubmit = async () => {
  const valid = await userFormRef.value.validate().catch(() => false)
  if (!valid) return
  
  submitLoading.value = true
  try {
    if (isEdit.value) {
      await updateUser(userForm)
      ElMessage.success('更新成功')
    } else {
      await createUser(userForm)
      ElMessage.success('创建成功')
    }
    dialogVisible.value = false
    getList()
  } catch (error) {
    console.error('提交失败：', error)
  } finally {
    submitLoading.value = false
  }
}

/**
 * 对话框关闭
 */
const handleDialogClose = () => {
  userFormRef.value?.resetFields()
  Object.assign(userForm, {
    userId: null,
    userName: '',
    nickName: '',
    password: '',
    email: '',
    phonenumber: '',
    sex: '0',
    status: 0,
    deptId: 100,
    remark: ''
  })
}

// 组件挂载时获取列表
onMounted(() => {
  getList()
})
</script>

<style lang="scss" scoped>
.user-container {
  .search-form {
    margin-bottom: 20px;
  }
  
  .toolbar {
    margin-bottom: 20px;
  }
  
  .pagination {
    margin-top: 20px;
    justify-content: flex-end;
  }
}
</style>
