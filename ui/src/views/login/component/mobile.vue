<template>
	<el-form size="large" class="login-content-form">
		<el-form-item class="login-animation1">
			<el-input text maxlength="11" :placeholder="$t('message.mobile.placeholder1')" v-model="state.ruleForm.userName" clearable autocomplete="off">
				<template #prefix>
					<i class="iconfont icon-dianhua el-input__icon"></i>
				</template>
			</el-input>
		</el-form-item>
		<el-form-item class="login-animation2">
			<el-col :span="15">
				<el-input text maxlength="4" :placeholder="$t('message.mobile.placeholder2')" v-model="state.ruleForm.code" clearable autocomplete="off">
					<template #prefix>
						<el-icon class="el-input__icon"><ele-Position /></el-icon>
					</template>
				</el-input>
			</el-col>
			<el-col :span="1"></el-col>
			<el-col :span="8">
        <el-button v-if="state.showCode" @click="sendCode" v-waves class="login-content-code">{{ state.codeText }}</el-button>
        <el-button v-else  v-waves class="login-content-code">{{ state.codeText }}</el-button>
			</el-col>
		</el-form-item>
		<el-form-item class="login-animation3">
			<el-button :loading="state.loading.signIn" @click="login" round type="primary" v-waves class="login-content-submit">
				<span>{{ $t('message.mobile.btnText') }}</span>
			</el-button>
		</el-form-item>
		<div class="font12 mt30 login-animation4 login-msg">{{ $t('message.mobile.msgText') }}</div>
	</el-form>
</template>

<script setup lang="ts" name="loginMobile">
import {computed, reactive} from 'vue';
import {signInCodeApi} from "/@/api/login";
import {ElMessage} from "element-plus";
import {initFrontEndControlRoutes} from "/@/router/frontEnd";
import {initBackEndControlRoutes} from "/@/router/backEnd";
import {Session} from "/@/utils/storage";
import {NextLoading} from "/@/utils/loading";
import {storeToRefs} from "pinia";
import {useThemeConfig} from "/@/stores/themeConfig";
import {useRoute, useRouter} from "vue-router";
import {formatAxis} from "/@/utils/formatTime";
import {useI18n} from "vue-i18n";
import Cookies from "js-cookie";
const storesThemeConfig = useThemeConfig();
const { themeConfig } = storeToRefs(storesThemeConfig);
const route = useRoute();
const router = useRouter();
// 定义变量内容
const { t } = useI18n();
// 定义变量内容
const state = reactive({
	ruleForm: {
		userName: '',
		code: '',
	},
  showCode:true,
  codeTimeNum:60,
  setIntervalTime:null,
  codeText:"获取验证码",
  loading: {
    signIn: false,
  },
});
const login = () => {
  if(!state.ruleForm.userName){
    ElMessage.error("手机号不能为空！")
    return
  }
  if(!state.ruleForm.code){
    ElMessage.error("验证码不能为空！")
    return
  }
  state.loading.signIn = true;
  signInCodeApi({
    userName:state.ruleForm.userName,
    code:state.ruleForm.code,
    loginType:"pc"
  }).then(res=>{
    if(res&&res.code==200){
      // ElMessage.success("登录成功！")
      //外部登录
      if(route.query.redirect_uri){
        window.location.href = route.query.redirect_uri;
      }else{
        // 存储 token 到浏览器缓存
        Session.set('token', res.token);
        // 模拟数据，对接接口时，记得删除多余代码及对应依赖的引入。用于 `/src/stores/userInfo.ts` 中不同用户登录判断（模拟数据）
        Cookies.set('userName', state.ruleForm.userName);
        signInSuccess()
      }
    }else{
      ElMessage.error("登录失败！")
      state.loading.signIn = false;
    }
  })
}
// 时间获取
const currentTime = computed(() => {
  return formatAxis(new Date());
});
// 登录成功后的跳转
const signInSuccess = async () => {
  let isNoPower=null
  if (!themeConfig.value.isRequestRoutes) {
    // 前端控制路由，2、请注意执行顺序
    isNoPower = await initFrontEndControlRoutes();
  } else {
    // 模拟后端控制路由，isRequestRoutes 为 true，则开启后端控制路由
    // 添加完动态路由，再进行 router 跳转，否则可能报错 No match found for location with path "/"
    isNoPower = await initBackEndControlRoutes();
  }
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
  state.loading.signIn = false;
}

const sendCode = () => {
  state.showCode=false
  state.codeTimeNum = 60
  state.setIntervalTime = setInterval(countDown, 1000)
}
const countDown= () => {
  // console.log(state.codeTimeNum)
  state.codeTimeNum--;
  if (state.codeTimeNum <= 0) {
    state.showCode=true
    state.codeText = "获取验证码";
    clearInterval(state.setIntervalTime);
  } else {
    state.codeText = state.codeTimeNum + "s"
  }
}
</script>

<style scoped lang="scss">
.login-content-form {
	margin-top: 20px;
	@for $i from 1 through 4 {
		.login-animation#{$i} {
			opacity: 0;
			animation-name: error-num;
			animation-duration: 0.5s;
			animation-fill-mode: forwards;
			animation-delay: calc($i/10) + s;
		}
	}
	.login-content-code {
		width: 100%;
		padding: 0;
	}
	.login-content-submit {
		width: 100%;
		letter-spacing: 2px;
		font-weight: 300;
		margin-top: 15px;
	}
	.login-msg {
		color: var(--el-text-color-placeholder);
	}
}
</style>
