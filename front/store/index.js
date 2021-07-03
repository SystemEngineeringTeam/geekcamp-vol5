export const state = () => ({
  text: '',
  count: 0,
  pastLatLng: { lat: 35.0036648, lng: 136.958297 },
  totalLength: 0,
  targetLength: 0,
})

export const getters = {
  getCount(state) {
    return state.count
  },
  getPastLatLng(state) {
    return state.pastLatLng
  },
  getTotalLength(state) {
    return state.totalLength
  },
  getTargetLength() {},
}

export const mutations = {
  increment(state) {
    state.count += 1
    console.log(state.count)
  },
  changText(state, text) {
    state.text = text
  },
  // ローカルストレージに保存
  saveTotalLength(state, length) {
    console.log(length)
    state.totalLength = length
  },
  savePosition(state, positon) {
    state.pastLatLng.lat = positon.lat
    state.pastLatLng.lng = positon.lng
  },
}

export const actions = {}
