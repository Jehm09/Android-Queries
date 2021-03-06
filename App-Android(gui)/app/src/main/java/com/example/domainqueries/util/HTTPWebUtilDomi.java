package com.example.domainqueries.util;

import android.util.Log;

import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.io.InputStream;
import java.net.HttpURLConnection;
import java.net.URL;


public class HTTPWebUtilDomi {

    private OnResponseListener listener;

    public HTTPWebUtilDomi() {

    }

    public void setListener(OnResponseListener listener) {
        this.listener = listener;
    }

    public void GETrequest(int callbackID, String url) {
        try {
            URL page = new URL(url);
            Log.e(">>>>>>>>>", url);
            HttpURLConnection connection = (HttpURLConnection) page.openConnection();
            connection.setRequestMethod("GET");
            InputStream is = connection.getInputStream();
            ByteArrayOutputStream baos = new ByteArrayOutputStream();
            byte[] buffer = new byte[4096];
            int bytesRead;
            while ((bytesRead = is.read(buffer)) != -1) {
                baos.write(buffer, 0, bytesRead);
            }
            is.close();
            baos.close();
            connection.disconnect();
            String response = new String(baos.toByteArray(), "UTF-8");
            Log.e(">>>>>>", response);
            if (listener != null) listener.onResponse(callbackID, response);
        }catch (IOException ex){
            ex.printStackTrace();
        }
    }


    public interface OnResponseListener {
        void onResponse(int callbackID, String response);
    }
}