export function pad(str, len, ch) {
    ch = ch || "0"
    let n = str.length
    while (n++ < len) {
        str = ch + str
    }
    return str
}

export function formatTime(time) {
    let str = pad(Math.floor(time / 3600).toString(), 2) + ":"
    time %= 3600
    return str + pad(Math.floor(time / 60).toString(), 2) + ":" + pad((time % 60).toString(), 2)
}

export function parseTime(time) {
    var arr = time.split(":")
    if (arr.length == 3) {
        return parseInt(arr[0]) * 3600 + parseInt(arr[1]) * 60 + parseInt(arr[2])
    }
    return 0
}