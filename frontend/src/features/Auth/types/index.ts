export interface SignUpRequest {
  email: string;
  password: string;
  username: string;
}

export interface SignUpResponse {
  id: number;
}

export interface SignUpVerifyRequest {
  code: string;
  id: number;
}

export interface SignInRequest {
  password: string;
  username: string;
}

export interface CheckTokenRequest {
  refresh_token: string;
}

export interface CheckTokenResponse {
  accessToken: string;
  refreshToken: string;
}
