<template>
    <ElForm>
        <ElFormItem label="视频：">
            <ElInput placeholder="请输入视频路径" v-model="video" />
        </ElFormItem>
        <ElFormItem label="设备：">
            <ElSelect v-model="device" placeholder="请选择设备" :no-data-text="selectTip" @visible-change="onVisibleChange">
                <ElOption v-for="value, index in devices" :value="index" :label="value" :key="index" />
            </ElSelect>
        </ElFormItem>
    </ElForm>
    <div>

    </div>
    <div class="horizontal">
        <ElText>{{ formatTime(time) }}</ElText>
        <ElSlider class="progress" v-model="time" :max="duration" :disabled="duration == 0"
            @change="onPositionChange" />
        <ElText>{{ formatTime(duration) }}</ElText>
    </div>
    <div class="horizontal">
        <ElButton @click="onStopClick" circle>
            <img class="icon" src="@/assets/stop.png">
        </ElButton>
        <ElButton @click="onPlayClick" circle size="large">
            <img v-if="paused || !playing" class="icon" src="@/assets/play.png">
            <img v-else class="icon" src="@/assets/pause.png">
        </ElButton>
        <ElButton @click="onMuteClick" circle>
            <img v-if="muted" class="icon" src="@/assets/mute.png">
            <img v-else class="icon" src="@/assets/volume.png">
        </ElButton>
    </div>
</template>

<script setup>
import { onMounted, onUnmounted, ref } from 'vue';
import { ElButton, ElForm, ElFormItem, ElInput, ElMessage, ElOption, ElSelect, ElSlider, ElText } from 'element-plus';
import request from './request';
import { formatTime, parseTime } from './utils/time';

const video = ref("")
const devices = ref([])
const device = ref()
const selectTip = ref("")
const playing = ref(false)
const paused = ref(false)
const muted = ref(false)
const volume = ref(0)
const time = ref(0)
const duration = ref(0)

function onVisibleChange(visible) {
    if (visible) {
        selectTip.value = "正在搜索投屏设备"
        devices.value = []
        request("/dlna/devices", {}, result => {
            if (result.devices.length > 0) {
                devices.value = result.devices
            } else {
                selectTip.value = "未搜索到可投屏设备"
            }
        })
    }
}

function onStopClick() {
    device.value = undefined
    playing.value = false
    paused.value = false
    time.value = 0
    duration.value = 0
    request("/dlna/close", {})
}

function onPlayClick() {
    if (playing.value) {
        if (paused.value) {
            request("/dlna/resume", {}, () => {
                paused.value = false
            })
        } else {
            request("/dlna/pause", {}, () => {
                paused.value = true
            })
        }
    } else {
        if (!video.value) {
            ElMessage("请输入视频路径")
            return
        }
        if (device.value === undefined) {
            ElMessage("请选择投屏设备")
            return
        }
        request("/dlna/play", { device: device.value, video: video.value }, () => {
            playing.value = true
            request("/dlna/mute/get", {}, result => {
                muted.value = result.muted
            })
            request("/dlna/volume/get", {}, result => {
                volume.value = result.volume
            })
        })
    }
}

function onMuteClick() {
    if (playing.value) {
        muted.value = !muted.value
        request("/dlna/mute/set", { mute: muted.value })
    }
}

function onPositionChange() {
    request("/dlna/seek", { time: formatTime(time.value) }, () => {
        paused.value = false
    })
}

function onVolumeChange() {
    request("/dlna/volume/set", { volume: volume.value })
}

let timer = 0;

function onKeydown(e) {
    switch (e.code) {
        case "ArrowLeft":
            time.value = Math.max(0, time.value - 5)
            ElMessage("快退5秒")
            onPositionChange()
            break
        case "ArrowRight":
            time.value = Math.min(duration.value, time.value + 5)
            ElMessage("快进5秒")
            onPositionChange()
            break
        case "ArrowUp":
            volume.value = Math.min(100, volume.value + 5)
            ElMessage("音量" + volume.value)
            onVolumeChange()
            break
        case "ArrowDown":
            volume.value = Math.max(0, volume.value - 5)
            ElMessage("音量" + volume.value)
            onVolumeChange()
            break
    }
}

onMounted(() => {
    timer = setInterval(() => {
        if (playing.value && !paused.value) {
            request("/dlna/position", {}, result => {
                time.value = parseTime(result.time)
                duration.value = parseTime(result.duration)
            })
        }
    }, 500);

    document.addEventListener("keydown", onKeydown)
})

onUnmounted(() => {
    clearInterval(timer)
    document.removeEventListener("keydown", onKeydown)
})

</script>

<style scoped>
.horizontal {
    display: flex;
    align-items: center;
    justify-content: center;
}

.progress {
    margin: 0 14px;
}

.icon {
    width: 20px;
    height: 20px;
}
</style>