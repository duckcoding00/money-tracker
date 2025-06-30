import { fail, type Actions } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load = (async () => {
	return {};
}) satisfies PageServerLoad;

export const actions: Actions = {
	reset: async ({ fetch, request }) => {
		const form = await request.formData();
		const username = form.get('username') as string;

		const body = JSON.stringify({
			username
		});

		if (username === '') {
			return fail(400, { message: 'DONT FORGET YOUR USERNAME' });
		}

		console.log(username);
		try {
			const response = await fetch('http://127.0.0.1:8080/api/v1/token/reset', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
					Accept: 'application/json'
				},
				body: body
			});

			const result = await response.json();

			return {
				result,
				username
			};
		} catch (error) {
			return fail(500, { message: 'Server error occurred' });
		}
	}
};
