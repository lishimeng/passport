<template>
  <el-form size="large" class="login-content-form" :rules="state.rules" ref="ruleForm" :model="state.form">
    <el-form-item prop="name">
      <el-input placeholder="请输入用户名" v-model="state.form.name"></el-input>
    </el-form-item>
    <el-form-item prop="password">
      <el-input placeholder="请输入密码" v-model="state.form.password" show-password></el-input>
    </el-form-item>
    <el-form-item prop="email">
      <el-input placeholder="请输入邮箱" v-model="state.form.email"></el-input>
    </el-form-item>
    <el-form-item prop="mobile">
      <el-input placeholder="请输入手机号" maxlength="11" v-model="state.form.mobile"></el-input>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" round @click="submitForm()">立即创建</el-button>
      <el-button round @click="resetForm()">重置</el-button>
    </el-form-item>
  </el-form>
</template>

<script setup lang="ts" name="register">
import {reactive, ref} from "vue";
import {registerApi} from "/@/api/register";
import {ElMessage} from "element-plus";

const ruleForm = ref()
const state = reactive({
  form: {
    name: "",
    password: "",
    email: "",
    mobile: "",
    verificationCode: ""
  },
  rules: {
    name: [
      {required: true, message: '请输入用户名', trigger: 'blur'},
    ],
    password: [
      {required: true, message: '请输入密码', trigger: 'blur'},
      {min: 6, message: "密码不能低于6位", trigger: "blur"},
    ],
    email: [
      {required: true, message: '请输入邮箱', trigger: 'blur'},
      {type: 'email', message: '请输入正确的邮箱', trigger: 'blur'}
    ],
    mobile: [
      {required: true, message: '请输入手机号', trigger: 'blur'},
      {
        required: true,
        pattern: /^1(3[0-9]|4[01456879]|5[0-35-9]|6[2567]|7[0-8]|8[0-9]|9[0-35-9])\d{8}$/,
        message: '请输入正确的手机号码',
        trigger: 'blur',
      },
    ]
  }
})
const submitForm = () => {
  ruleForm.value.validate((valid: any) => {
    if (valid) {
      console.log('submit!!');
      registerApi(state.form).then(res => {
        if (res && res.code == 200) {
          ElMessage.success("注册成功！")
        } else {
          ElMessage.error(res.message + "！")
        }
      })
    } else {
      console.log('error submit!!');
      return false;
    }
  });
}
const resetForm = () => {
  ruleForm.value.resetFields()
}
</script>

<style scoped>

</style>