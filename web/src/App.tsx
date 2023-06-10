import { useState, useEffect } from 'react';
import './App.css';
import { Button } from './components/button';
import { HTTPMethod, fetcher } from './utils/helpers';

function App() {
	const [ws, setWs] = useState<WebSocket | null>(null);

	const pingAPI = async () => {
		const path = 'ping';
		try {
			const response = await fetcher<{ message: string }>(path, HTTPMethod.GET);
			console.log(response.message);
		} catch (error) {
			console.log(error);
		}
	};

	useEffect(() => {
		console.log('useeffect called');
		if (!ws) {
			console.log('creating websocket....');
			const ws = new WebSocket(
				'ws://localhost:8080/connect?username=sebastian&channelId=12345'
			);
			console.log(ws);
			setWs(ws);

			ws.onopen = () => {
				console.log('websocket connection opened');
			};
			ws.onclose = () => {
				console.log('websocket connection closed');
			};
			ws.onerror = (err) => {
				console.log('websocket connection error', err);
			};
			ws.onmessage = (msg) => {
				console.log('websocket message received', msg);
			};
		}
		// pingAPI();
	}, [ws]);

	const handleSendMessage = () => {
		if (!ws || ws.readyState != ws.OPEN) return;
		console.log('Sending message, the ws is defined and ready');
		const message = {
			type: 'chatMsg',
			data: {
				text: 'Hello from the client',
			},
		};
		ws.send(JSON.stringify(message));
	};

	const handleDisconnect = () => {
		if (!ws || ws.readyState != ws.OPEN) return;
		console.log('Disconnecting');
		ws.close();
	};

	return (
		<>
			<div className="flex items-center gap-5 p-5 border rounded-lg">
				<Button variant="outline" onClick={handleSendMessage}>
					Send message
				</Button>
				<Button variant="outline" onClick={handleDisconnect}>
					Disconnect
				</Button>
			</div>
		</>
	);
}

export default App;
