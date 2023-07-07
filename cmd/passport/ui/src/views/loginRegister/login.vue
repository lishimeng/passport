<template>
  <el-container>
    <el-header style="background-color: #ffffff;display: flex;align-items: center;">
      <img style="width: 30px;height: 30px;display: block;" src="../../assets/c1-register.svg"/>
      <span>Thingple-IOT</span>
    </el-header>
    <el-main>
      <el-card class="box-card">
        <template #header>
          <div style="text-align: center;">
            <h2>用户登录</h2>
          </div>
        </template>
        <div>
          <el-form :rules="state.rules" ref="ruleForm" :model="state.form" label-width="120px">
            <el-form-item label="用户名" prop="userName">
              <el-input placeholder="用户名/手机号/邮箱" v-model="state.form.userName"></el-input>
            </el-form-item>
            <el-form-item label="密码" prop="password">
              <el-input placeholder="请输入密码" v-model="state.form.password" show-password></el-input>
            </el-form-item>
            <el-form-item label="验证码" prop="code">
              <el-input placeholder="请输入验证码" v-model="state.form.code">
                <template #append>
                  <span v-if="state.showCode==true" @click="codeTimeOut()">{{ state.showCodeText }}</span>
                  <span v-else>{{ state.showCodeText }}</span>
                </template>
              </el-input>
            </el-form-item>
            <el-form-item>
              <el-button :loading="state.loading" type="primary" @click="submitForm()">立即登录</el-button>
              <el-button type="primary" @click="toPage()">注册</el-button>
            </el-form-item>
          </el-form>
        </div>
      </el-card>
    </el-main>
  </el-container>
</template>

<script setup lang="ts" name="login">
import {computed, onMounted, reactive, ref} from "vue";
import {useRoute, useRouter} from 'vue-router';
import {Session} from "/@/utils/storage";
import {useI18n} from "vue-i18n";
import {ElMessage} from "element-plus";
import {formatAxis} from "/@/utils/formatTime";
import {useThemeConfig} from "/@/stores/themeConfig";
import {storeToRefs} from "pinia";
import {NextLoading} from "/@/utils/loading";
import {signInApi} from "/@/api/login";

const {t} = useI18n();
const storesThemeConfig = useThemeConfig();
const {themeConfig} = storeToRefs(storesThemeConfig);
const route = useRoute();
const router = useRouter();
const ruleForm = ref()
const state = reactive({
  isShowPassword: false,
  form: {
    userName: "",
    password: "",
    code: ""
  },
  rules: {
    userName: [
      {required: true, message: '请输入用户名', trigger: 'blur'},
    ],
    password: [
      {required: true, message: '请输入密码', trigger: 'blur'},
    ],
    code: [
      {required: true, message: '请输入验证码', trigger: 'blur'},
    ],
  },
  showCode: true,
  showCodeText: "获取验证码",
  setIntervalTime: null,
  timeNum: 60,
  loading: false,
})
const codeTimeOut = () => {
  state.showCode = false;
  state.timeNum = 60
  state.setIntervalTime = setInterval(timeNum, 1000)
}
const timeNum = () => {
  console.log(state.timeNum)
  state.timeNum--;
  if (state.timeNum <= 0) {
    state.showCode = true;
    state.showCodeText = "获取验证码";
    clearInterval(state.setIntervalTime);
  } else {
    state.showCodeText = state.timeNum + "s"
  }
}
onMounted(()=>{
  console.log("s_url:",route.query.redirect_uri)
})
// 时间获取
const currentTime = computed(() => {
  return formatAxis(new Date());
});
const submitForm = () => {
  ruleForm.value.validate((valid: any) => {
    if (valid) {
      console.log('submit!!');
      signInApi(state.form).then(res => {
        if (res && res.code == 200) {
          console.log(res)
          ElMessage.success('登录成功！');
          window.location.href = route.query.redirect_uri;
        }
      })
      /*state.loading = true;
      // 存储 token 到浏览器缓存
      Session.set('token', Math.random().toString(36).substr(0));
      // 模拟数据，对接接口时，记得删除多余代码及对应依赖的引入。用于 `/src/stores/userInfo.ts` 中不同用户登录判断（模拟数据）
      Cookies.set('userName', state.form.userName);
      if (!themeConfig.value.isRequestRoutes) {
        // 前端控制路由，2、请注意执行顺序
        const isNoPower = await initFrontEndControlRoutes();
        signInSuccess(isNoPower);
      } else {
        // 模拟后端控制路由，isRequestRoutes 为 true，则开启后端控制路由
        // 添加完动态路由，再进行 router 跳转，否则可能报错 No match found for location with path "/"
        const isNoPower = await initBackEndControlRoutes();
        // 执行完 initBackEndControlRoutes，再执行 signInSuccess
        signInSuccess(isNoPower);
      }*/
    } else {
      console.log('error submit!!');
      return false;
    }
  });
}
// 登录成功后的跳转
const signInSuccess = (isNoPower: boolean | undefined) => {
  if (isNoPower) {
    ElMessage.warning('抱歉，您没有登录权限');
    Session.clear();
  } else {
    // 初始化登录成功时间问候语
    let currentTimeInfo = currentTime.value;
    // 登录成功，跳到转首页
    // 如果是复制粘贴的路径，非首页/登录页，那么登录成功后重定向到对应的路径中
    if (route.query?.redirect) {
      router.push({
        path: <string>route.query?.redirect,
        query: Object.keys(<string>route.query?.params).length > 0 ? JSON.parse(<string>route.query?.params) : '',
      });
    } else {
      router.push('/');
    }
    // 登录成功提示
    const signInText = t('message.signInText');
    ElMessage.success(`${currentTimeInfo}，${signInText}`);
    // 添加 loading，防止第一次进入界面时出现短暂空白
    NextLoading.start();
  }
  state.loading = false;
}
const toPage = () => {
  router.push({
    path: "/register"
  })
}
</script>

<style scoped>

</style>