// ignore_for_file: public_member_api_docs, sort_constructors_first
class ViagemModel {
  String name;
  int id;
  String dataChegada;
  String dataSaida;
  int valor;
  ViagemModel({
    required this.name,
    required this.id,
    required this.dataChegada,
    required this.dataSaida,
    required this.valor,
  });

  Map<String, dynamic> toMap() {
    return <String, dynamic>{
      'name': name,
      'id': id,
      'data_chegada': dataChegada,
      'data_saida': dataSaida,
      'valor': valor
    };
  }

  factory ViagemModel.fromJson(Map<String, dynamic> json) {
    return ViagemModel(
      name: json['name'] as String,
      id: json['id'] as int,
      dataChegada: json['data_chegada'] as String,
      dataSaida: json['data_saida'] as String,
      valor: json['valor'] as int,
    );
  }
}
