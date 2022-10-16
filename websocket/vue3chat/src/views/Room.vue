<template>
	<div class="home">
		<div class="main_box">
			<RoomItem
				:class="{ reverse: item.is_self }"
				v-for="(item, index) in message_list"
				:key="index"
				:reverse="item.is_self ? true : false"
				:content="item"
			></RoomItem>
			<el-empty v-if="message_list.length == 0" description="暂无记录" />
			<div class="bottom_container">
				<div class="input_container">
					<input
						type="text"
						placeholder="请输入内容"
						@keyup.enter="submitMsg"
						v-model="message_info.content"
					/>
					<button @click="submitMsg">发送</button>
				</div>
			</div>
		</div>
	</div>
</template>

<script setup>
	import RoomItem from "@/components/RoomItem.vue";
	import { onMounted, ref, reactive, nextTick } from "vue";
	import WebsocketHeartbeatJs from "websocket-heartbeat-js";
	import { useRoute } from "vue-router";
	import { useUserStore } from "@/store/index";
	import { ElMessage } from "element-plus";
	import { getMessageList } from "@/api/index";
	let route = useRoute();
	let userStore = useUserStore();
	const message_info = reactive({
		content: "",
		group: "",
		uid: userStore.uid,
		user_name: userStore.user_name,
		avatar: userStore.avatar,
	});
	message_info.group = route.query.room;
	let message_list = reactive([]);
	let ws = new WebsocketHeartbeatJs({
		url: "ws://192.168.31.44:9999/socket/" + route.query.room,
		// url: "ws://192.168.1.140:9999/socket",
	});
	ws.onopen = function () {
		console.log("connect success");
		ws.send("hello server");
	};
	ws.onmessage = function (e) {
		try {
			let res = JSON.parse(e.data);
			let obj = {};
			if (res.status == 6) {
				let is_self = userStore.uid != res.data.uid ? false : true;
				obj = {
					content: res.data.content,
					group: res.data.group,
					time: res.data.time,
					uid: res.data.uid,
					user_name: res.data.user_name,
					is_self: is_self,
					avatar: res.data.avatar,
				};
				message_list.push(obj);
				nextTick(() => {
					var height = document.body.scrollHeight;
					window.scroll({ top: height, left: 0, behavior: "smooth" });
				});
			}
		} catch (error) {}
	};
	ws.onreconnect = function () {
		console.log("reconnecting...");
	};
	onMounted(() => {
		setTimeout(() => {}, 2000);
	});

	function submitMsg() {
		if (message_info.content == "") {
			return;
		}
		message_info.uid = userStore.uid;
		message_info.user_name = userStore.user_name;
		message_info.avatar = userStore.avatar;
		if (!message_info.uid) {
			ElMessage({
				message: "请稍后 获取uid失败",
				type: "warning",
			});
			return;
		}
		let data = {
			status: 6,
			data: Object.assign({}, message_info),
		};
		message_info.content = "";
		let message = JSON.stringify(data);
		ws.send(message);
	}
	const getMessage = async () => {
		let parmas = {
			group: route.query.room,
			page: 1,
			size: 100,
		};
		let res = await getMessageList(parmas);
		if (res.code === 200 && res.data.data) {
			let list = [];
			let obj = {};
			res.data.data.map((e) => {
				let is_self = userStore.uid != e.u_id ? false : true;
				console.log(userStore.uid, e.u_id);
				obj = {
					content: e.content,
					group: e.group,
					time: e.c_time,
					uid: e.u_id,
					user_name: e.account,
					is_self: is_self,
					avatar: e.avatar,
				};
				message_list.push(obj);
			});
		}
		nextTick(() => {
			var height = document.body.scrollHeight;
			window.scroll({ top: height, left: 0, behavior: "smooth" });
		});
	};
	getMessage();
</script>


<style lang="less" scoped>
.home {
	display: flex;
	justify-content: center;
	position: relative;
	background-image: radial-gradient(
		circle,
		#d5eae7,
		#d7ebe8,
		#d9ede9,
		#dbeeea,
		#ddefeb,
		#dbeeea,
		#daede9,
		#d8ece8,
		#d2e9e4,
		#cbe6e1,
		#c5e2dd,
		#bedfda
	);
	min-height: calc(100vh - 50px);
}
.main_box {
	width: 900px;
	min-height: calc(100vh - 50px);
	background-color: rgba(255, 255, 255, 0.5);
	display: flex;
	flex-direction: column;
	padding: 50px 0 100px 0;
	box-sizing: border-box;
}
.reverse {
	align-self: flex-end;
}

.bottom_container {
	width: 900px;
	height: 100px;
	position: fixed;
	bottom: 0;
	left: 50%;
	transform: translateX(-50%);
	box-sizing: border-box;
	padding: 15px;
	background-color: #edf7f5;
	border-top: 2px solid #d8ece8;
}
.input_container {
	display: flex;
	align-items: center;
	height: 30px;
	padding: 10px;
	border-radius: 5px;
	background-color: #fff;
}
.input_container input {
	flex: 1;
	height: 100%;
	box-sizing: border-box;
	outline: none;
	padding: 0 10px;
	border: 1px solid #acd6d0;
}
.input_container button {
	display: block;
	width: 70px;
	height: 100%;
	margin-left: 10px;
	border: 1px solid #acd6d0;
	background-color: #dbeeea;
	color: #3c7a71;
}
.input_container button:active {
	background-color: #c3e1dd;
	color: #3c7a71;
}
</style>
