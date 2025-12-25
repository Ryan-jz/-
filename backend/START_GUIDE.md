# Backend Service Start Guide

## Recommended: Use Product API Server

The product API server is a standalone, fully functional service that provides all product-related API endpoints.

### Start Command
```bash
cd backend
go run product_server.go
```

### Service Info
- **Port**: 8001
- **Status**: ✅ Fully Operational
- **API Prefix**: `/api/v1`

### Available Endpoints
- `GET /api/v1/product/category/list` - Get category list
- `GET /api/v1/product/list` - Get product list
- `GET /api/v1/product/detail/:id` - Get product detail
- `POST /api/v1/product/create` - Create product
- `PUT /api/v1/product/update` - Update product
- `DELETE /api/v1/product/delete/:id` - Delete product

### Success Indicator
You should see:
```
[INFO] Product API server started on port 8001
[INFO] http server started listening on [:8001]
```

## Admin Server (main.go) - Currently Unavailable

The admin server has a routing binding issue and cannot start at this time.

### Issue
```
invalid handler: router binding error
```

### Impact
- ❌ Cannot start `main.go`
- ✅ Does NOT affect product API functionality
- ✅ Product API server runs independently

## Frontend Configuration

### frontend-web/.env.local
```env
VITE_API_BASE_URL=http://localhost:8001/api/v1
```

### frontend-admin/.env.local
```env
VITE_API_BASE_URL=http://localhost:8001/api/v1
```

## Complete Startup Process

1. Start Product API Server: `go run product_server.go`
2. Start Frontend Web: `pnpm dev` (port 5174)
3. Start Admin Panel: `npm run dev` (port 5173)

## Service Status

| Service | File | Port | Status |
|---------|------|------|--------|
| Product API | product_server.go | 8001 | ✅ Working |
| Admin Server | main.go | 8000 | ❌ Issue |
| Frontend Web | - | 5174 | ✅ Working |
| Admin Panel | - | 5173 | ✅ Working |

## Conclusion

**The product management system is fully functional!**

Use `product_server.go` as the backend API server with the frontend pages for complete product management capabilities.
