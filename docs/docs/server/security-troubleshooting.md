# Security & Troubleshooting

## Security Considerations

- **Network Binding**: The server runs on all interfaces (`0.0.0.0`) by default. In production environments, consider binding to specific interfaces using the `--host` flag.

- **Authentication**: There is no built-in authentication mechanism. For production deployments, consider adding a reverse proxy with authentication if needed.

- **Data Storage**: Test executions are stored in memory only. Consider implementing persistent storage solutions for production use cases.

- **CORS**: Cross-Origin Resource Sharing (CORS) is enabled by default. You can disable it using `--cors=false` if not needed.

## Troubleshooting

### Server Won't Start

**Problem**: Server fails to start or reports port binding errors.

**Solutions**:

1. Check if the port is already in use:
   ```bash
   lsof -i :8080
   ```

2. Try a different port:
   ```bash
   voiceflow server --port 9090
   ```

3. Check if you have permission to bind to the port (especially for ports < 1024):
   ```bash
   sudo voiceflow server --port 80
   ```

### API Returns 404

**Problem**: API endpoints return 404 Not Found errors.

**Solutions**:

1. Ensure you're using the correct base path `/api/v1/` for API endpoints
2. Verify the server is running and accessible
3. Check the server logs for any startup errors

### Logs Not Appearing

**Problem**: Test execution logs are not visible or incomplete.

**Solutions**:

1. Enable debug mode to see more detailed logging:
   ```bash
   voiceflow server --debug
   ```

2. Check that the test suite path is correct and accessible
3. Verify that environment variables (VF_API_KEY, etc.) are properly set

### Connection Refused Errors

**Problem**: Cannot connect to the server from external clients.

**Solutions**:

1. Verify the server is bound to the correct interface:
   ```bash
   voiceflow server --host 0.0.0.0 --port 8080
   ```

2. Check firewall settings and ensure the port is open
3. For local testing, try connecting to `127.0.0.1` instead of `localhost`

### High Memory Usage

**Problem**: Server consumes excessive memory during long-running operations.

**Solutions**:

1. Monitor test execution status and clean up completed executions
2. Consider implementing execution cleanup routines
3. Restart the server periodically for long-running deployments

### Swagger Documentation Not Loading

**Problem**: Cannot access Swagger UI at `/swagger/index.html`.

**Solutions**:

1. Ensure Swagger is enabled (it's enabled by default):
   ```bash
   voiceflow server --swagger=true
   ```

2. Try accessing the full URL: `http://localhost:8080/swagger/index.html`
3. Check browser console for JavaScript errors
4. Verify CORS settings if accessing from a different domain
