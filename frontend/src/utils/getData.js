const SERVER_URL = "http://localhost:5500"

export async function getDomain(url) {
    const resp = await fetch(`${SERVER_URL}/info/${url}`)
    return resp.json()
}

export async function getHistory() {
    const resp = await fetch(`${SERVER_URL}/history`)
    return resp.json()
}