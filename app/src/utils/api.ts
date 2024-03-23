const baseUrl ='http://localhost:8080/'

const api = {
  async post(path: string, body: never) {
    const response = await fetch(`${baseUrl}${path}`, {
      method: 'POST',
      body
    })
    return response.json()
  },
  async get(path: string) {
    const response = await fetch(`${baseUrl}${path}`)
    return response.json()
  },
  async put() { },
  async patch() { },
  async delete() {}
}



export default api;