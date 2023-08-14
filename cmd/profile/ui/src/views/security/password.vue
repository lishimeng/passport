<template>
	<div class="personal layout-pd">
		<el-row>
			<!-- 更新信息 -->
			<el-col :span="24">
				<el-card shadow="hover" class="mt15 personal-edit" header="MFA">
					<div class="personal-edit-title mb15">账号安全</div>
					<div class="personal-edit-safe-box">
						<div class="personal-edit-safe-item">
							<div class="personal-edit-safe-item-left">
								<div class="personal-edit-safe-item-left-label">修改密码</div>
								<div class="personal-edit-safe-item-left-value">密码：{{ mfaPasswordState.password }}</div>
							</div>
							<div class="personal-edit-safe-item-right">
								<el-button text type="primary" @click="showDrawer()">立即修改</el-button>
							</div>
						</div>
					</div>
				</el-card>
			</el-col>
		</el-row>
		<el-drawer title="修改密码" v-model="state.showDrawer">
			<el-form :model="mfaPasswordState.passwordForm" style="padding: 30px 30px;" label-width="100px">
				<el-form-item label="新的密码">
					<el-input v-model="mfaPasswordState.passwordForm.password" placeholder="请输入新的密码"></el-input>
				</el-form-item>
				<el-row justify="center">
					<el-col :span="6">
						<el-button type="primary" style="width: 200px" @click="submit()"> 确定 </el-button>
					</el-col>
				</el-row>
			</el-form>
		</el-drawer>
	</div>
</template>

<script setup lang="ts" name="MfaEmail">
import { reactive, computed } from 'vue';
import { formatAxis } from '/@/utils/formatTime';
import { storeToRefs } from "pinia";
import { ElMessage, ElMessageBox } from 'element-plus';
import { changePasswordApi } from '/@/api/login'
import { useUserInfo } from "/@/stores/userInfo";
import { logout } from '/@/utils/passport';
import { Local } from '/@/utils/storage';
const stores = useUserInfo();
const { userInfos } = storeToRefs(stores);
// 定义变量内容
const mfaPasswordState = reactive({
	password: "******",
	passwordForm: {
		password: '',
		name: userInfos.value.userName
	}
});
const state = reactive({
	showDrawer: false,
})

const showDrawer = () => {
	mfaPasswordState.passwordForm.password = ""
	state.showDrawer = true
}

const submit = () => {
	if (mfaPasswordState.passwordForm.password.length == 0) {
		ElMessage.warning("密码不能为空!")
		return
	}
	changePasswordApi({
		password: mfaPasswordState.passwordForm.password
	}).then((res) => {
		if (res && res.code == 200) {
			ElMessage.success("修改成功")
			state.showDrawer = false
			setTimeout(() => {
				ElMessageBox.confirm('修改成功是否返回登录', '提示', {
					confirmButtonText: '返回登录',
					cancelButtonText: '取消',
				}).then(() => {
					Local.remove("token")
					logout()
				}).catch(() => { });
			}, 1000)

		} else {
			ElMessage.error("修改失败! ")
		}
	}).catch((err) => {
		console.log(err)
	})
}

// 当前时间提示语
const currentTime = computed(() => {
	return formatAxis(new Date());
});
</script>

<style scoped lang="scss">
@import '../../theme/mixins/index.scss';

.personal {
	.personal-user {
		height: 130px;
		display: flex;
		align-items: center;

		.personal-user-left {
			width: 100px;
			height: 130px;
			border-radius: 3px;

			:deep(.el-upload) {
				height: 100%;
			}

			.personal-user-left-upload {
				img {
					width: 100%;
					height: 100%;
					border-radius: 3px;
				}

				&:hover {
					img {
						animation: logoAnimation 0.3s ease-in-out;
					}
				}
			}
		}

		.personal-user-right {
			flex: 1;
			padding: 0 15px;

			.personal-title {
				font-size: 18px;
				@include text-ellipsis(1);
			}

			.personal-item {
				display: flex;
				align-items: center;
				font-size: 13px;

				.personal-item-label {
					color: var(--el-text-color-secondary);
					@include text-ellipsis(1);
				}

				.personal-item-value {
					@include text-ellipsis(1);
				}
			}
		}
	}

	.personal-info {
		.personal-info-more {
			float: right;
			color: var(--el-text-color-secondary);
			font-size: 13px;

			&:hover {
				color: var(--el-color-primary);
				cursor: pointer;
			}
		}

		.personal-info-box {
			height: 130px;
			overflow: hidden;

			.personal-info-ul {
				list-style: none;

				.personal-info-li {
					font-size: 13px;
					padding-bottom: 10px;

					.personal-info-li-title {
						display: inline-block;
						@include text-ellipsis(1);
						color: var(--el-text-color-secondary);
						text-decoration: none;
					}

					& a:hover {
						color: var(--el-color-primary);
						cursor: pointer;
					}
				}
			}
		}
	}

	.personal-recommend-row {
		.personal-recommend-col {
			.personal-recommend {
				position: relative;
				height: 100px;
				border-radius: 3px;
				overflow: hidden;
				cursor: pointer;

				&:hover {
					i {
						right: 0px !important;
						bottom: 0px !important;
						transition: all ease 0.3s;
					}
				}

				i {
					position: absolute;
					right: -10px;
					bottom: -10px;
					font-size: 70px;
					transform: rotate(-30deg);
					transition: all ease 0.3s;
				}

				.personal-recommend-auto {
					padding: 15px;
					position: absolute;
					left: 0;
					top: 5%;
					color: var(--next-color-white);

					.personal-recommend-msg {
						font-size: 12px;
						margin-top: 10px;
					}
				}
			}
		}
	}

	.personal-edit {
		.personal-edit-title {
			position: relative;
			padding-left: 10px;
			color: var(--el-text-color-regular);

			&::after {
				content: '';
				width: 2px;
				height: 10px;
				position: absolute;
				left: 0;
				top: 50%;
				transform: translateY(-50%);
				background: var(--el-color-primary);
			}
		}

		.personal-edit-safe-box {
			border-bottom: 1px solid var(--el-border-color-light, #ebeef5);
			padding: 15px 0;

			.personal-edit-safe-item {
				width: 100%;
				display: flex;
				align-items: center;
				justify-content: space-between;

				.personal-edit-safe-item-left {
					flex: 1;
					overflow: hidden;

					.personal-edit-safe-item-left-label {
						color: var(--el-text-color-regular);
						margin-bottom: 5px;
					}

					.personal-edit-safe-item-left-value {
						color: var(--el-text-color-secondary);
						@include text-ellipsis(1);
						margin-right: 15px;
					}
				}
			}

			&:last-of-type {
				padding-bottom: 0;
				border-bottom: none;
			}
		}
	}
}</style>
