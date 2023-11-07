<template>
  <div class="form_container">
    <div class="title">上传视频</div>
    <Form @submit="onSubmit">
      <CellGroup inset class="cell">
        <div class="video_des">
          <Field
            name="uploader"
            class="uploader"
            :rules="[{ required: true, message: '请添加视频' }]"
          >
            <template #input>
              <Uploader
                v-model="videourl"
                :after-read="afterRead"
                :before-delete="deletevideo"
                accept="video/*"
                :max-count="1"
                required
              >
              </Uploader>
            </template>
          </Field>

          <Field
            class="des_text"
            v-model="description"
            name="description"
            maxlength="150"
            type="textarea"
            :autosize="{
              maxHeight: 140,
              minHeight: 140,
            }"
            :rules="[{ required: true, message: '请填写视频描述' }]"
            show-word-limit
            placeholder="添加作品描述"
          />
        </div>
        <Field
          v-model="location"
          is-link
          readonly
          name="area"
          label="你在哪里？"
          placeholder="点击添加位置"
          @click="showArea = true"
        />
        <Popup v-model:show="showArea" position="bottom">
          <Area :area-list="areaList" @confirm="onConfirm" @cancel="onCancel" />
        </Popup>
        <Field
          v-model="privacy"
          is-link
          readonly
          label="可见范围"
          name="privacy"
          placeholder="请选择可见范围"
          @click="showPrivacy = true"
        />
        <Popup v-model:show="showPrivacy" round position="bottom">
          <Cascader
            v-model="cascaderValue"
            title="请选择可见范围"
            :options="options"
            @close="showPrivacy = false"
            @finish="onFinish"
          />
        </Popup>
      </CellGroup>
      <div style="margin: 16px">
        <Button round block type="primary" native-type="submit"> 提交 </Button>
        <div style="margin: 16px"></div>
        <Button round block type="danger" @click="cancelSubmit"> 取消 </Button>
      </div>
    </Form>
  </div>
</template>
<script setup>
import { ref, reactive } from "vue";
import {
  Form,
  Field,
  CellGroup,
  Button,
  Uploader,
  Popup,
  Area,
  Cascader,
  ImagePreview,
} from "vant";
import { areaList } from "@vant/area-data";
import { useRouter } from "vue-router";
const fileUrl = ref("");
const description = ref("");
const videourl = reactive([]);
const images = reactive([]);
const onSubmit = (values) => {
  console.log("submit", values);
  // area
  // :
  // '北京市/北京市/西城区'
  // description
  // :
  // "123"
  // privacy
  // :
  // "公开"
  // uploader
  // :
  // [Proxy(Object)]
};
const afterRead = (file) => {
  // 此时可以自行将文件上传至服务器
  let binaryData = [];
  binaryData.push(file);

  let url = window.URL.createObjectURL(new Blob(binaryData));
  videourl.push({ url });
  // var formDate = new FormData();
  // formDate.append("file", file.file);
};
const deletevideo = () => {
  videourl.length = 0;
  console.log("@@@");
};
const location = ref("");
const showArea = ref(false);
const onConfirm = ({ selectedOptions }) => {
  showArea.value = false;
  location.value = selectedOptions.map((item) => item.text).join("/");
};
const onCancel = () => {
  showArea.value = false;
  location.value = "";
};
const showPrivacy = ref(false);
const privacy = ref("公开");
const cascaderValue = ref("");
// 选项列表，children 代表子选项，支持多级嵌套
const options = [
  {
    text: "公开",
    value: "0",
  },
  {
    text: "好友可见",
    value: "1",
  },
  {
    text: "自己可见",
    value: "2",
  },
];
// 全部选项选择完毕后，会触发 finish 事件
const onFinish = ({ selectedOptions }) => {
  showPrivacy.value = false;
  privacy.value = selectedOptions.map((option) => option.text).join("/");
};
const router = useRouter();
const cancelSubmit = () => {
  fileUrl.value = "";
  description.value = "";
  location.value = "";
  showArea.value = false;
  showPrivacy.value = false;
  privacy.value = "公开";
  cascaderValue.value = "";
  router.go(-1);
};
</script>
<style scoped>
.form_container {
  width: 100vw;
  height: 100vh;
  padding: 1rem 0;
  background-color: #161823;
}
.cell {
  height: 6.5rem;
  padding: 0.5rem 0;
}
.title {
  text-align: center;
  padding: 2rem auto;
  margin-bottom: 1.5rem;
  font-size: 0.5rem;
  color: #fff;
}
.video_des {
  display: flex;
  height: 3.8rem;
}
.uploader {
  padding: 0.2rem;
  width: 3.5rem;
  padding-right: 0;
  font-size: 1rem;
}
:deep(.van-uploader__upload) {
  width: 2rem;
  height: 3rem;
}
:deep(.van-cell__right-icon) {
  margin: var(--van-padding-base);
}
:deep(.van-uploader__preview-delete-icon) {
  top: -0.35rem;
  right: 0.35rem;
  transform: scale(0.2) translate(10%, -10%);
}
.des_text {
  width: 100%;
}
.preview-cover {
  position: absolute;
  bottom: 0;
  box-sizing: border-box;
  width: 100%;
  padding: 4px;
  color: #fff;
  font-size: 12px;
  text-align: center;
  background: rgba(0, 0, 0, 0.3);
}
</style>
