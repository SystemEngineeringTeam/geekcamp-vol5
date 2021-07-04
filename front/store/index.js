export const state = () => ({
  pastLatLng: { lat: 34.994697, lng: 136.969749 },
  totalLength: 0,

  MidTimeTmp: 0,
  targetLength: 0,
  target: 0,
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
  getTarget(state) {
    return state.target
  },
}

export const mutations = {
  increment(state) {
    state.count += 1
  },
  changText(state, text) {
    state.text = text
  },
  // ローカルストレージに保存
  saveTotalLength(state, length) {
    state.totalLength = length
  },
  savePosition(state, positon) {
    state.pastLatLng.lat = positon.lat
    state.pastLatLng.lng = positon.lng
  },
  saveMidTimeTmp(state, time) {
    console.log(time)
    state.MidTimeTmp = time
  },
  setTarget(state, target) {
    state.target = target
  },
  clear(state) {
    state.totalLength = 0
    state.target = 0
  },
}

export const actions = {}
