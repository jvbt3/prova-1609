import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:provaesoft8s/model/model.dart';

class ViagemService {
  Future<http.Response> postData(
      {required String name,
      required int id,
      required String dataChegada,
      required String dataSaida,
      required double valor}) async {
    try {
      final response = await http.post(
        Uri.parse('http://localhost:3000/api/viagem'),
        headers: <String, String>{
          'Content-Type': 'application/json; charset=UTF-8',
        },
        body: jsonEncode(<String, dynamic>{
          'name': name,
          'id': id,
          'data_chegada': dataChegada,
          'data_saida': dataSaida,
          'valor': valor
        }),
      );

      return response;
    } catch (e) {
      throw Exception('Error: $e');
    }
  }

  Future<List<ViagemModel>> fetchData() async {
    try {
      final request = await http.get(
        Uri.parse('http://localhost:3000/api/viagem'),
        headers: <String, String>{
          'Content-Type': 'application/json; charset=UTF-8',
        },
      );

      if (request.statusCode == 200) {
        final Map<String, dynamic> jsonBody = jsonDecode(request.body);
        var list = jsonBody['data'] as List;
        List<ViagemModel> viagens =
            list.map((u) => ViagemModel.fromJson(u)).toList();

        return viagens;
      } else {
        throw Exception('Failed to load user data');
      }
    } catch (e) {
      throw Exception('Erro na requisição $e');
    }
  }

  Future<void> updateData(
      {required int id,
      required String name,
      required String dataChegada,
      required String dataSaida,
      required double valor}) async {
    try {
      // ignore: unused_local_variable
      final response = await http.patch(
        Uri.parse('http://localhost:3000/api/viagem/$id'),
        headers: <String, String>{
          'Content-Type': 'application/json; charset=UTF-8',
        },
        body: jsonEncode(<String, dynamic>{
          'name': name,
          'data_saida': dataSaida,
          'data_chegada': dataChegada,
          'valor': valor
        }),
      );
    } catch (e) {
      throw Exception('Erro na requisição $e');
    }
  }

  Future<void> deleteData({required int id}) async {
    try {
      // ignore: unused_local_variable
      final response = await http.delete(
        Uri.parse('http://localhost:3000/api/viagem/$id'),
        headers: <String, String>{
          'Content-Type': 'application/json; charset=UTF-8',
        },
      );
    } catch (e) {
      throw Exception('Erro na requisição $e');
    }
  }
}
