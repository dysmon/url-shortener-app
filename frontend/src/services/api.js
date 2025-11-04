const API_BASE = process.env.REACT_APP_API_BASE || '';

export const shortenUrl = async (originalUrl) => {
  const response = await fetch(`${API_BASE}/shorten`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ url: originalUrl }),
  });

  if (!response.ok) {
    const error = await response.json();
    throw new Error(error.message || 'Failed to shorten URL');
  }

  const data = await response.json();
  return data.ShortURL;
};