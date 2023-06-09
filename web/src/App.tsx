import { useState, useEffect } from 'react';
import './App.css';
import { Button } from './components/button';

function App() {
	const [ws, setWs] = useState<WebSocket | null>(null);

	useEffect(() => {
		console.log('useeffect called');
		if (!ws) {
			console.log('creating websocket....');
			const ws = new WebSocket('ws://localhost:8080');
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
	}, [ws]);

	const handleConnect = () => {
		if (!ws) return;
		ws.send('connect');
	};

	return (
		<>
			<div>
				<Button variant="outline" onClick={handleConnect}>
					connect
				</Button>
			</div>
		</>
	);
}

export default App;
