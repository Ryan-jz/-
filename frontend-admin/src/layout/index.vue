<!--
  主布局组件
  包含侧边栏、顶部导航栏和内容区域
-->
<template>
  <el-container class="layout-container">
    <!-- 侧边栏 -->
    <el-aside :width="isCollapse ? '64px' : '200px'" class="layout-aside">
      <div class="logo">
        <span v-if="!isCollapse">GF Admin</span>
        <span v-else>GF</span>
      </div>
      
      <!-- 菜单 -->
      <el-menu
        :default-active="activeMenu"
        :collapse="isCollapse"
        :unique-opened="true"
        router
        background-color="#304156"
        text-color="#bfcbd9"
        active-text-color="#409eff"
      >
        <el-menu-item index="/dashboard">
          <el-icon><HomeFilled /></el-icon>
          <template #title>首页</template>
        </el-menu-item>
        
        <el-sub-menu index="/product">
          <template #title>
            <el-icon><Goods /></el-icon>
            <span>产品管理</span>
          </template>
          <el-menu-item index="/product/list">
            <el-icon><List /></el-icon>
            <template #title>产品列表</template>
          </el-menu-item>
          <el-menu-item index="/product/category">
            <el-icon><Menu /></el-icon>
            <template #title>分类管理</template>
          </el-menu-item>
        </el-sub-menu>
        
        <el-menu-item index="/recipe/list">
          <el-icon><Reading /></el-icon>
          <template #title>食谱管理</template>
        </el-menu-item>
        
        <el-sub-menu index="/system">
          <template #title>
            <el-icon><Setting /></el-icon>
            <span>系统管理</span>
          </template>
          <el-menu-item index="/system/user">
            <el-icon><User /></el-icon>
            <template #title>用户管理</template>
          </el-menu-item>
          <el-menu-item index="/system/role">
            <el-icon><UserFilled /></el-icon>
            <template #title>角色管理</template>
          </el-menu-item>
          <el-menu-item index="/system/menu">
            <el-icon><Menu /></el-icon>
            <template #title>菜单管理</template>
          </el-menu-item>
        </el-sub-menu>
      </el-menu>
    </el-aside>

    <!-- 主内容区 -->
    <el-container>
      <!-- 顶部导航栏 -->
      <el-header class="layout-header">
        <div class="header-left">
          <el-icon class="collapse-icon" @click="toggleCollapse">
            <Fold v-if="!isCollapse" />
            <Expand v-else />
          </el-icon>
        </div>
        
        <div class="header-right">
          <el-dropdown @command="handleCommand">
            <span class="user-info">
              <el-avatar :size="32" :src="userInfo?.avatar" />
              <span class="username">{{ userInfo?.nickName || '管理员' }}</span>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item command="profile">个人中心</el-dropdown-item>
                <el-dropdown-item command="logout" divided>退出登录</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </el-header>

      <!-- 内容区域 -->
      <el-main class="layout-main">
        <router-view />
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useUserStore } from '@/stores/user'
import { ElMessageBox } from 'element-plus'

const route = useRoute()
const userStore = useUserStore()

// 侧边栏折叠状态
const isCollapse = ref(false)

// 当前激活的菜单
const activeMenu = computed(() => route.path)

// 用户信息
const userInfo = computed(() => userStore.userInfo)

// 切换侧边栏折叠状态
const toggleCollapse = () => {
  isCollapse.value = !isCollapse.value
}

// 下拉菜单命令处理
const handleCommand = (command) => {
  if (command === 'logout') {
    ElMessageBox.confirm('确定要退出登录吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      userStore.logout()
    })
  } else if (command === 'profile') {
    // 跳转到个人中心
    console.log('个人中心')
  }
}

// 组件挂载时获取用户信息
onMounted(() => {
  if (!userInfo.value) {
    userStore.getUserInfo()
  }
})
</script>

<style lang="scss" scoped>
.layout-container {
  height: 100%;
}

.layout-aside {
  background-color: #304156;
  transition: width 0.3s;
  
  .logo {
    height: 60px;
    line-height: 60px;
    text-align: center;
    font-size: 20px;
    font-weight: bold;
    color: #fff;
    background-color: #2b3a4b;
  }
  
  .el-menu {
    border-right: none;
  }
}

.layout-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: #fff;
  border-bottom: 1px solid #e6e6e6;
  padding: 0 20px;
  
  .header-left {
    .collapse-icon {
      font-size: 20px;
      cursor: pointer;
      
      &:hover {
        color: #409eff;
      }
    }
  }
  
  .header-right {
    .user-info {
      display: flex;
      align-items: center;
      cursor: pointer;
      
      .username {
        margin-left: 10px;
      }
    }
  }
}

.layout-main {
  background-color: #f0f2f5;
  padding: 20px;
}
</style>
