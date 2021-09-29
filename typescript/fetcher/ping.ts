import { Empty } from 'assets/protobuf/web/ping_pb'
import { PingServiceClient } from 'assets/protobuf/web/PingServiceClientPb'

const hostname = "https://test-api.dev.snm.djty.co"
const client = new PingServiceClient(hostname, {}, { withCredentials: true })

export const GetNow = async (): Promise<number> => {
    const request = new Empty()
    try {
        const t = await client.now(request, {})
        return t.getT()
    } catch (e) {
        throw e
    }
}
