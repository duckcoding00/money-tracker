import { fail, isRedirect, redirect, type HandleFetch } from '@sveltejs/kit';
import { jwtDecode } from 'jwt-decode';

export const handleFetch: HandleFetch = async ({ event, request, fetch }) => {
	const cookie = event.cookies.get('access_token');

	if (request.url.startsWith('http://127.0.0.1:8080/')) {
		let currentToken = cookie;

		if (cookie) {
			try {
				const decode = jwtDecode(cookie);
				const currentTime = Math.floor(Date.now() / 1000);

				if (decode.exp && decode.exp < currentTime + 300) {
					console.log('Token expired or will expire soon, refreshing...');

					try {
						const refreshResponse = await fetch('http://127.0.0.1:8080/api/v1/token/refresh', {
							method: 'POST',
							headers: {
								'Content-Type': 'application/json',
								Authorization: `Bearer ${cookie}`
							}
						});

						if (!refreshResponse.ok) {
							event.cookies.delete('access_token', { path: '/' });
							throw redirect(303, '/login');
						}

						const result = await refreshResponse.json();
						const newAccessToken = result.data?.access_token;

						if (!newAccessToken) {
							throw new Error('No access token in refresh response');
						}

						event.cookies.set('access_token', newAccessToken, {
							path: '/',
							httpOnly: true,
							sameSite: 'lax',
							secure: false,
							maxAge: 60 * 60 * 24 * 7
						});

						currentToken = newAccessToken;
						console.log('Token refreshed successfully');
					} catch (refreshError) {
						console.error('Token refresh failed:', refreshError);
						event.cookies.delete('access_token', { path: '/' });

						if (isRedirect(refreshError)) throw refreshError;
						throw redirect(303, '/login');
					}
				}
			} catch (error) {
				console.error('JWT decode error:', error);
				event.cookies.delete('access_token', { path: '/' });

				if (isRedirect(error)) throw error;
				throw redirect(303, '/login');
			}
		}
		if (currentToken) {
			request.headers.set('Authorization', `Bearer ${currentToken}`);
		}

		if (!request.headers.get('Content-Type') && request.method !== 'GET') {
			request.headers.set('Content-Type', 'application/json');
		}
	}

	try {
		const response = await fetch(request);
		if (response.status === 401 && cookie) {
			console.log('Received 401, clearing token');
			event.cookies.delete('access_token', { path: '/' });
		}

		return response;
	} catch (error) {
		console.error('Fetch error:', error);
		throw error;
	}
};
