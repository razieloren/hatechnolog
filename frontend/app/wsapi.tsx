import {Websocket, WebsocketBuilder} from 'websocket-ts';
import {messages as wrapper} from './messages/wrapper'

export function InitAPIWS(endpoint: wrapper.Endpoint, onReady?: (instance: Websocket) => void, onMessage?: (instance: Websocket, data: any) => void): Websocket {
    //const builder = new WebsocketBuilder(process.env.API_WS_ADDRESS!)
    const builder = new WebsocketBuilder("ws://localhost:8080/ws")
    if (onMessage !== undefined) {
        builder.onMessage((ws, e) => {
            const reader = new FileReader();
            reader.onload = () => {
                onMessage(ws, reader.result);
            };
            reader.readAsArrayBuffer(e.data);
        })
    }
    builder.onOpen((innerWs, _) => {
        const endpoointIntent = new wrapper.Wrapper({
            endpoint_intent: new wrapper.EndpointIntent({
                endpoint: endpoint
            })
        });
        innerWs.send(endpoointIntent.serialize());
        if (onReady !== undefined) {
            onReady(innerWs);
        }
    })
    const ws = builder.build()
    return ws
}
