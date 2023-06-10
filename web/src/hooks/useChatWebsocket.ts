import { useCallback, useEffect, useRef, useState } from 'react';

export const useChatWebsocket = (channelId: string) => {
	const [data, setData] = useState<string>('');
	const [error, setError] = useState<string>('');
	const websocketConnection = useRef<WebSocket | null>(null);

	useEffect(() => {
		const token = localStorage.getItem('token');
		if (!token) {
			setError('Unauthorized');
			return;
		}

		const ws = new WebSocket(
			`ws://localhost:8080/connect?channelId=${channelId}&token=${token}`
		);
		websocketConnection.current = ws;

		ws.onmessage = (event) => {
			const data = event.data;
			setData(data);
		};

		return () => {
			ws.close();
		};
	}, [channelId]);

	// add a new type for message and more neccessary fields
	const send = useCallback((message: string) => {
		if (!websocketConnection.current) {
			return;
		}
		const data = {
			type: 'chatMsg',
			data: {
				message: message,
			},
		};

		websocketConnection.current.send(JSON.stringify(data));
	}, []);

	return { data, error, send };
};
