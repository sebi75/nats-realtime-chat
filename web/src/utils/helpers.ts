export enum HTTPMethod {
	GET = 'GET',
	POST = 'POST',
	PUT = 'PUT',
	DELETE = 'DELETE',
}

export const fetcher = async <T>(
	path: string,
	method: HTTPMethod,
	params?: Record<string, string>
): Promise<T> => {
	const baseURL = getServerURL();
	const url = `${baseURL}/${path}`;

	try {
		const response = await fetch(url, {
			headers: {
				'Content-Type': 'application/json',
				Accept: 'application/json',
				Authorization: `Bearer ${localStorage.getItem('token')}`,
			},
			method: method,
			body: JSON.stringify(params),
		});

		if (!response.ok) {
			throw new Error(response.statusText);
		}

		const data = await response.json();
		return data;
	} catch (error: any) {
		throw new Error(error.message || 'Something went wrong');
	}
};

export const getServerURL = () => {
	return 'http://localhost:8080';
};

export const restService = {
	get: async (path: string) => {
		return fetcher(path, HTTPMethod.GET);
	},
};
