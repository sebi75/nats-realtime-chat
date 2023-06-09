export enum HTTPMethod {
	GET = 'GET',
	POST = 'POST',
	PUT = 'PUT',
	DELETE = 'DELETE',
}

export const fetcher = async (
	path: string,
	method: HTTPMethod,
	params?: Record<string, string>
) => {
	const baseURL = getURLFromEnv();
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
	} catch (error: any) {
		throw new Error(error.message || 'Something went wrong');
	}
};

export const getURLFromEnv = () => {
	return process.env.REACT_APP_API_URL;
};

export const restService = {
	get: async (path: string) => {
		return fetcher(path, HTTPMethod.GET);
	},
};
