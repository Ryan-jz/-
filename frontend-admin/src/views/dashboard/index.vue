<!--
  首页仪表盘
  展示系统概览和统计数据
-->
<template>
  <div class="dashboard-container">
    <el-row :gutter="20">
      <!-- 统计卡片 -->
      <el-col :span="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-icon" style="background-color: #409eff;">
              <el-icon :size="30"><User /></el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">1,234</div>
              <div class="stat-label">用户总数</div>
            </div>
          </div>
        </el-card>
      </el-col>
      
      <el-col :span="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-icon" style="background-color: #67c23a;">
              <el-icon :size="30"><Document /></el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">5,678</div>
              <div class="stat-label">文章总数</div>
            </div>
          </div>
        </el-card>
      </el-col>
      
      <el-col :span="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-icon" style="background-color: #e6a23c;">
              <el-icon :size="30"><View /></el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">98,765</div>
              <div class="stat-label">访问量</div>
            </div>
          </div>
        </el-card>
      </el-col>
      
      <el-col :span="6">
        <el-card class="stat-card">
          <div class="stat-content">
            <div class="stat-icon" style="background-color: #f56c6c;">
              <el-icon :size="30"><Message /></el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">432</div>
              <div class="stat-label">消息数</div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 欢迎信息 -->
    <el-card class="welcome-card" style="margin-top: 20px;">
      <template #header>
        <div class="card-header">
          <span>欢迎使用 GF Admin 管理系统</span>
        </div>
      </template>
      <div class="welcome-content">
        <h3>系统介绍</h3>
        <p>这是一个基于 GoFrame v2 + Vue3 + Element Plus 开发的后台管理系统。</p>
        <h3 style="margin-top: 20px;">技术栈</h3>
        <ul>
          <li>后端：GoFrame v2.6.4 + MySQL 8.0 + Redis</li>
          <li>前端：Vue 3 + Vite + Element Plus + Pinia</li>
          <li>认证：JWT Token</li>
        </ul>
        <h3 style="margin-top: 20px;">功能特性</h3>
        <ul>
          <li>用户管理：用户的增删改查、权限分配</li>
          <li>角色管理：角色的增删改查、权限配置</li>
          <li>菜单管理：动态菜单配置</li>
          <li>部门管理：组织架构管理</li>
        </ul>
      </div>
    </el-card>
  </div>
</template>

<script setup>
// 首页组件
</script>

<style lang="scss" scoped>
.dashboard-container {
  .stat-card {
    .stat-content {
      display: flex;
      align-items: center;
      
      .stat-icon {
        width: 60px;
        height: 60px;
        border-radius: 10px;
        display: flex;
        align-items: center;
        justify-content: center;
        color: #fff;
        margin-right: 20px;
      }
      
      .stat-info {
        flex: 1;
        
        .stat-value {
          font-size: 24px;
          font-weight: bold;
          color: #333;
          margin-bottom: 5px;
        }
        
        .stat-label {
          font-size: 14px;
          color: #999;
        }
      }
    }
  }
  
  .welcome-card {
    .welcome-content {
      h3 {
        color: #333;
        margin-bottom: 10px;
      }
      
      p {
        color: #666;
        line-height: 1.8;
      }
      
      ul {
        padding-left: 20px;
        
        li {
          color: #666;
          line-height: 2;
        }
      }
    }
  }
}
</style>
